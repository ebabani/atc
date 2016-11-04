// This file was generated by counterfeiter
package workerfakes

import (
	"io"
	"sync"

	"github.com/concourse/atc/worker"
)

type FakeImageFetchingDelegate struct {
	StderrStub        func() io.Writer
	stderrMutex       sync.RWMutex
	stderrArgsForCall []struct{}
	stderrReturns     struct {
		result1 io.Writer
	}
	ImageVersionDeterminedStub        func(worker.ResourceCacheIdentifier) error
	imageVersionDeterminedMutex       sync.RWMutex
	imageVersionDeterminedArgsForCall []struct {
		arg1 worker.ResourceCacheIdentifier
	}
	imageVersionDeterminedReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImageFetchingDelegate) Stderr() io.Writer {
	fake.stderrMutex.Lock()
	fake.stderrArgsForCall = append(fake.stderrArgsForCall, struct{}{})
	fake.recordInvocation("Stderr", []interface{}{})
	fake.stderrMutex.Unlock()
	if fake.StderrStub != nil {
		return fake.StderrStub()
	} else {
		return fake.stderrReturns.result1
	}
}

func (fake *FakeImageFetchingDelegate) StderrCallCount() int {
	fake.stderrMutex.RLock()
	defer fake.stderrMutex.RUnlock()
	return len(fake.stderrArgsForCall)
}

func (fake *FakeImageFetchingDelegate) StderrReturns(result1 io.Writer) {
	fake.StderrStub = nil
	fake.stderrReturns = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakeImageFetchingDelegate) ImageVersionDetermined(arg1 worker.ResourceCacheIdentifier) error {
	fake.imageVersionDeterminedMutex.Lock()
	fake.imageVersionDeterminedArgsForCall = append(fake.imageVersionDeterminedArgsForCall, struct {
		arg1 worker.ResourceCacheIdentifier
	}{arg1})
	fake.recordInvocation("ImageVersionDetermined", []interface{}{arg1})
	fake.imageVersionDeterminedMutex.Unlock()
	if fake.ImageVersionDeterminedStub != nil {
		return fake.ImageVersionDeterminedStub(arg1)
	} else {
		return fake.imageVersionDeterminedReturns.result1
	}
}

func (fake *FakeImageFetchingDelegate) ImageVersionDeterminedCallCount() int {
	fake.imageVersionDeterminedMutex.RLock()
	defer fake.imageVersionDeterminedMutex.RUnlock()
	return len(fake.imageVersionDeterminedArgsForCall)
}

func (fake *FakeImageFetchingDelegate) ImageVersionDeterminedArgsForCall(i int) worker.ResourceCacheIdentifier {
	fake.imageVersionDeterminedMutex.RLock()
	defer fake.imageVersionDeterminedMutex.RUnlock()
	return fake.imageVersionDeterminedArgsForCall[i].arg1
}

func (fake *FakeImageFetchingDelegate) ImageVersionDeterminedReturns(result1 error) {
	fake.ImageVersionDeterminedStub = nil
	fake.imageVersionDeterminedReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageFetchingDelegate) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.stderrMutex.RLock()
	defer fake.stderrMutex.RUnlock()
	fake.imageVersionDeterminedMutex.RLock()
	defer fake.imageVersionDeterminedMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeImageFetchingDelegate) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ worker.ImageFetchingDelegate = new(FakeImageFetchingDelegate)
