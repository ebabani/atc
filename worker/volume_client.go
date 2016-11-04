package worker

import (
	"errors"
	"time"

	"code.cloudfoundry.org/clock"
	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/dbng"
	"github.com/concourse/baggageclaim"
)

const creatingVolumeRetryDelay = 1 * time.Second

//go:generate counterfeiter . VolumeClient

type VolumeClient interface {
	FindOrCreateVolumeForResourceCache(
		lager.Logger,
		VolumeSpec,
		*dbng.UsedResourceCache,
	) (Volume, error)
	FindOrCreateVolumeForContainer(
		lager.Logger,
		VolumeSpec,
		*dbng.CreatingContainer,
		*dbng.Team,
		string,
	) (Volume, error)
	FindOrCreateVolumeForBaseResourceType(
		lager.Logger,
		VolumeSpec,
		*dbng.Team,
		string,
	) (Volume, error)
	ListVolumes(lager.Logger, VolumeProperties) ([]Volume, error)
	LookupVolume(lager.Logger, string) (Volume, bool, error)
}

var ErrVolumeExpiredImmediately = errors.New("volume expired immediately after saving")
var ErrCreatedVolumeNotFound = errors.New("failed-to-find-created-volume-in-baggageclaim")
var ErrBaseResourceTypeNotFound = errors.New("base-resource-type-not-found")

type volumeClient struct {
	baggageclaimClient        baggageclaim.Client
	db                        GardenWorkerDB
	volumeFactory             VolumeFactory
	dbVolumeFactory           dbng.VolumeFactory
	dbBaseResourceTypeFactory dbng.BaseResourceTypeFactory
	clock                     clock.Clock
	dbWorker                  *dbng.Worker
}

func NewVolumeClient(
	baggageclaimClient baggageclaim.Client,
	db GardenWorkerDB,
	volumeFactory VolumeFactory,
	dbVolumeFactory dbng.VolumeFactory,
	dbBaseResourceTypeFactory dbng.BaseResourceTypeFactory,
	clock clock.Clock,
	dbWorker *dbng.Worker,
) VolumeClient {
	return &volumeClient{
		baggageclaimClient:        baggageclaimClient,
		db:                        db,
		volumeFactory:             volumeFactory,
		dbVolumeFactory:           dbVolumeFactory,
		dbBaseResourceTypeFactory: dbBaseResourceTypeFactory,
		clock:    clock,
		dbWorker: dbWorker,
	}
}

func (c *volumeClient) FindOrCreateVolumeForContainer(
	logger lager.Logger,
	volumeSpec VolumeSpec,
	container *dbng.CreatingContainer,
	team *dbng.Team,
	mountPath string,
) (Volume, error) {
	return c.findOrCreateVolume(
		logger,
		volumeSpec,
		func() (dbng.CreatingVolume, dbng.CreatedVolume, error) {
			return c.dbVolumeFactory.FindContainerVolume(team, c.dbWorker, container, mountPath)
		},
		func() (dbng.CreatingVolume, error) {
			return c.dbVolumeFactory.CreateContainerVolume(team, c.dbWorker, container, mountPath)
		},
	)
}

func (c *volumeClient) FindOrCreateVolumeForBaseResourceType(
	logger lager.Logger,
	volumeSpec VolumeSpec,
	team *dbng.Team,
	resourceTypeName string,
) (Volume, error) {
	baseResourceType, found, err := c.dbBaseResourceTypeFactory.Find(resourceTypeName)
	if err != nil {
		return nil, err
	}
	if !found {
		logger.Error("base-resource-type-not-found", ErrBaseResourceTypeNotFound, lager.Data{"resource-type-name": resourceTypeName})
		return nil, ErrBaseResourceTypeNotFound
	}

	return c.findOrCreateVolume(
		logger,
		volumeSpec,
		func() (dbng.CreatingVolume, dbng.CreatedVolume, error) {
			return c.dbVolumeFactory.FindBaseResourceTypeVolume(team, c.dbWorker, baseResourceType)
		},
		func() (dbng.CreatingVolume, error) {
			return c.dbVolumeFactory.CreateBaseResourceTypeVolume(team, c.dbWorker, baseResourceType)
		},
	)
}

func (c *volumeClient) FindOrCreateVolumeForResourceCache(
	logger lager.Logger,
	volumeSpec VolumeSpec,
	usedResourceCache *dbng.UsedResourceCache,
) (Volume, error) {
	return c.findOrCreateVolume(
		logger,
		volumeSpec,
		func() (dbng.CreatingVolume, dbng.CreatedVolume, error) {
			return c.dbVolumeFactory.FindResourceCacheVolume(c.dbWorker, usedResourceCache)
		},
		func() (dbng.CreatingVolume, error) {
			return c.dbVolumeFactory.CreateResourceCacheVolume(c.dbWorker, usedResourceCache)
		},
	)
}

func (c *volumeClient) ListVolumes(logger lager.Logger, properties VolumeProperties) ([]Volume, error) {
	if c.baggageclaimClient == nil {
		return []Volume{}, nil
	}

	bcVolumes, err := c.baggageclaimClient.ListVolumes(
		logger,
		baggageclaim.VolumeProperties(properties),
	)
	if err != nil {
		logger.Error("failed-to-list-volumes", err)
		return nil, err
	}

	volumes := []Volume{}
	for _, bcVolume := range bcVolumes {
		volume, err := c.volumeFactory.BuildWithIndefiniteTTL(logger, bcVolume)
		if err != nil {
			return []Volume{}, err
		}

		volumes = append(volumes, volume)
	}

	return volumes, nil
}

func (c *volumeClient) LookupVolume(logger lager.Logger, handle string) (Volume, bool, error) {
	if c.baggageclaimClient == nil {
		return nil, false, nil
	}

	bcVolume, found, err := c.baggageclaimClient.LookupVolume(logger, handle)
	if err != nil {
		logger.Error("failed-to-lookup-volume", err)
		return nil, false, err
	}

	if !found {
		return nil, false, nil
	}

	volume, err := c.volumeFactory.BuildWithIndefiniteTTL(logger, bcVolume)
	if err != nil {
		return nil, false, err
	}

	return volume, true, nil
}

func (c *volumeClient) selectLowestAlphabeticalVolume(logger lager.Logger, volumes []db.SavedVolume) (db.SavedVolume, error) {
	var lowestVolume db.SavedVolume

	for _, v := range volumes {
		if lowestVolume.ID == 0 {
			lowestVolume = v
		} else if v.ID < lowestVolume.ID {
			lowestVolume = v
		}
	}

	for _, v := range volumes {
		if v != lowestVolume {
			expLog := logger.Session("expiring-redundant-volume", lager.Data{
				"volume-handle": v.Handle,
			})

			err := c.expireVolume(expLog, v.Handle)
			if err != nil {
				return db.SavedVolume{}, err
			}
		}
	}

	return lowestVolume, nil
}

func (c *volumeClient) expireVolume(logger lager.Logger, handle string) error { // TODO consider removing this method?
	logger.Info("expiring")

	wVol, found, err := c.LookupVolume(logger, handle)
	if err != nil {
		logger.Error("failed-to-look-up-volume", err)
		return err
	}

	if !found {
		logger.Debug("volume-already-gone")
		return nil
	}

	logger.Debug("releasing a volume " + handle + " [super logs]")

	wVol.Destroy()

	return nil
}

func (c *volumeClient) findOrCreateVolume(
	logger lager.Logger,
	volumeSpec VolumeSpec,
	findVolumeFunc func() (dbng.CreatingVolume, dbng.CreatedVolume, error),
	createVolumeFunc func() (dbng.CreatingVolume, error),
) (Volume, error) {
	var bcVolume baggageclaim.Volume
	var bcVolumeFound bool

	creatingVolume, createdVolume, err := findVolumeFunc()
	if err != nil {
		logger.Error("failed-to-find-volume-in-db", err)
		return nil, err
	}

	if createdVolume != nil {
		bcVolume, bcVolumeFound, err = c.baggageclaimClient.LookupVolume(
			logger.Session("lookup-volume"),
			createdVolume.Handle(),
		)
		if err != nil {
			logger.Error("failed-to-lookup-volume-in-baggageclaim", err)
			return nil, err
		}

		if !bcVolumeFound {
			logger.Error("failed-to-lookup-volume-in-baggageclaim", ErrCreatedVolumeNotFound)
			return nil, ErrCreatedVolumeNotFound
		}
	} else {
		if creatingVolume != nil {
			bcVolume, bcVolumeFound, err = c.baggageclaimClient.LookupVolume(
				logger.Session("create-volume"),
				creatingVolume.Handle(),
			)
			if err != nil {
				logger.Error("failed-to-lookup-volume-in-baggageclaim", err)
				return nil, err
			}
		} else {
			creatingVolume, err = createVolumeFunc()
			if err != nil {
				logger.Error("failed-to-create-volume-in-db", err)
				return nil, err
			}
		}

		lock, acquired, err := c.db.AcquireVolumeCreatingLock(logger, creatingVolume.ID())
		if err != nil {
			logger.Error("failed-to-acquire-volume-creating-lock", err)
			return nil, err
		}

		if !acquired {
			c.clock.Sleep(creatingVolumeRetryDelay)
			return c.findOrCreateVolume(logger, volumeSpec, findVolumeFunc, createVolumeFunc)
		}

		defer lock.Release()

		if !bcVolumeFound {
			bcVolume, err = c.baggageclaimClient.CreateVolume(
				logger.Session("create-volume"),
				creatingVolume.Handle(),
				volumeSpec.baggageclaimVolumeSpec(),
			)
			if err != nil {
				logger.Error("failed-to-create-volume-in-baggageclaim", err)
				return nil, err
			}
		}

		_, err = creatingVolume.Created()
		if err != nil {
			logger.Error("failed-to-initialize-volume", err)
			return nil, err
		}
	}

	volume, err := c.volumeFactory.BuildWithIndefiniteTTL(logger, bcVolume)
	if err != nil {
		logger.Error("failed-build-volume", err)
		return nil, err
	}

	return volume, nil
}
