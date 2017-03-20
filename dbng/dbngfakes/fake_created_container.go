// This file was generated by counterfeiter
package dbngfakes

import (
	"sync"

	"github.com/concourse/atc/dbng"
)

type FakeCreatedContainer struct {
	IDStub        func() int
	iDMutex       sync.RWMutex
	iDArgsForCall []struct{}
	iDReturns     struct {
		result1 int
	}
	iDReturnsOnCall map[int]struct {
		result1 int
	}
	HandleStub        func() string
	handleMutex       sync.RWMutex
	handleArgsForCall []struct{}
	handleReturns     struct {
		result1 string
	}
	handleReturnsOnCall map[int]struct {
		result1 string
	}
	DiscontinueStub        func() (dbng.DestroyingContainer, error)
	discontinueMutex       sync.RWMutex
	discontinueArgsForCall []struct{}
	discontinueReturns     struct {
		result1 dbng.DestroyingContainer
		result2 error
	}
	discontinueReturnsOnCall map[int]struct {
		result1 dbng.DestroyingContainer
		result2 error
	}
	DestroyingStub        func() (dbng.DestroyingContainer, error)
	destroyingMutex       sync.RWMutex
	destroyingArgsForCall []struct{}
	destroyingReturns     struct {
		result1 dbng.DestroyingContainer
		result2 error
	}
	destroyingReturnsOnCall map[int]struct {
		result1 dbng.DestroyingContainer
		result2 error
	}
	IsHijackedStub        func() bool
	isHijackedMutex       sync.RWMutex
	isHijackedArgsForCall []struct{}
	isHijackedReturns     struct {
		result1 bool
	}
	isHijackedReturnsOnCall map[int]struct {
		result1 bool
	}
	WorkerNameStub        func() string
	workerNameMutex       sync.RWMutex
	workerNameArgsForCall []struct{}
	workerNameReturns     struct {
		result1 string
	}
	workerNameReturnsOnCall map[int]struct {
		result1 string
	}
	MarkAsHijackedStub        func() error
	markAsHijackedMutex       sync.RWMutex
	markAsHijackedArgsForCall []struct{}
	markAsHijackedReturns     struct {
		result1 error
	}
	markAsHijackedReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCreatedContainer) ID() int {
	fake.iDMutex.Lock()
	ret, specificReturn := fake.iDReturnsOnCall[len(fake.iDArgsForCall)]
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct{}{})
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if fake.IDStub != nil {
		return fake.IDStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.iDReturns.result1
}

func (fake *FakeCreatedContainer) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeCreatedContainer) IDReturns(result1 int) {
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeCreatedContainer) IDReturnsOnCall(i int, result1 int) {
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeCreatedContainer) Handle() string {
	fake.handleMutex.Lock()
	ret, specificReturn := fake.handleReturnsOnCall[len(fake.handleArgsForCall)]
	fake.handleArgsForCall = append(fake.handleArgsForCall, struct{}{})
	fake.recordInvocation("Handle", []interface{}{})
	fake.handleMutex.Unlock()
	if fake.HandleStub != nil {
		return fake.HandleStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.handleReturns.result1
}

func (fake *FakeCreatedContainer) HandleCallCount() int {
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	return len(fake.handleArgsForCall)
}

func (fake *FakeCreatedContainer) HandleReturns(result1 string) {
	fake.HandleStub = nil
	fake.handleReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreatedContainer) HandleReturnsOnCall(i int, result1 string) {
	fake.HandleStub = nil
	if fake.handleReturnsOnCall == nil {
		fake.handleReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.handleReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreatedContainer) Discontinue() (dbng.DestroyingContainer, error) {
	fake.discontinueMutex.Lock()
	ret, specificReturn := fake.discontinueReturnsOnCall[len(fake.discontinueArgsForCall)]
	fake.discontinueArgsForCall = append(fake.discontinueArgsForCall, struct{}{})
	fake.recordInvocation("Discontinue", []interface{}{})
	fake.discontinueMutex.Unlock()
	if fake.DiscontinueStub != nil {
		return fake.DiscontinueStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.discontinueReturns.result1, fake.discontinueReturns.result2
}

func (fake *FakeCreatedContainer) DiscontinueCallCount() int {
	fake.discontinueMutex.RLock()
	defer fake.discontinueMutex.RUnlock()
	return len(fake.discontinueArgsForCall)
}

func (fake *FakeCreatedContainer) DiscontinueReturns(result1 dbng.DestroyingContainer, result2 error) {
	fake.DiscontinueStub = nil
	fake.discontinueReturns = struct {
		result1 dbng.DestroyingContainer
		result2 error
	}{result1, result2}
}

func (fake *FakeCreatedContainer) DiscontinueReturnsOnCall(i int, result1 dbng.DestroyingContainer, result2 error) {
	fake.DiscontinueStub = nil
	if fake.discontinueReturnsOnCall == nil {
		fake.discontinueReturnsOnCall = make(map[int]struct {
			result1 dbng.DestroyingContainer
			result2 error
		})
	}
	fake.discontinueReturnsOnCall[i] = struct {
		result1 dbng.DestroyingContainer
		result2 error
	}{result1, result2}
}

func (fake *FakeCreatedContainer) Destroying() (dbng.DestroyingContainer, error) {
	fake.destroyingMutex.Lock()
	ret, specificReturn := fake.destroyingReturnsOnCall[len(fake.destroyingArgsForCall)]
	fake.destroyingArgsForCall = append(fake.destroyingArgsForCall, struct{}{})
	fake.recordInvocation("Destroying", []interface{}{})
	fake.destroyingMutex.Unlock()
	if fake.DestroyingStub != nil {
		return fake.DestroyingStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.destroyingReturns.result1, fake.destroyingReturns.result2
}

func (fake *FakeCreatedContainer) DestroyingCallCount() int {
	fake.destroyingMutex.RLock()
	defer fake.destroyingMutex.RUnlock()
	return len(fake.destroyingArgsForCall)
}

func (fake *FakeCreatedContainer) DestroyingReturns(result1 dbng.DestroyingContainer, result2 error) {
	fake.DestroyingStub = nil
	fake.destroyingReturns = struct {
		result1 dbng.DestroyingContainer
		result2 error
	}{result1, result2}
}

func (fake *FakeCreatedContainer) DestroyingReturnsOnCall(i int, result1 dbng.DestroyingContainer, result2 error) {
	fake.DestroyingStub = nil
	if fake.destroyingReturnsOnCall == nil {
		fake.destroyingReturnsOnCall = make(map[int]struct {
			result1 dbng.DestroyingContainer
			result2 error
		})
	}
	fake.destroyingReturnsOnCall[i] = struct {
		result1 dbng.DestroyingContainer
		result2 error
	}{result1, result2}
}

func (fake *FakeCreatedContainer) IsHijacked() bool {
	fake.isHijackedMutex.Lock()
	ret, specificReturn := fake.isHijackedReturnsOnCall[len(fake.isHijackedArgsForCall)]
	fake.isHijackedArgsForCall = append(fake.isHijackedArgsForCall, struct{}{})
	fake.recordInvocation("IsHijacked", []interface{}{})
	fake.isHijackedMutex.Unlock()
	if fake.IsHijackedStub != nil {
		return fake.IsHijackedStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isHijackedReturns.result1
}

func (fake *FakeCreatedContainer) IsHijackedCallCount() int {
	fake.isHijackedMutex.RLock()
	defer fake.isHijackedMutex.RUnlock()
	return len(fake.isHijackedArgsForCall)
}

func (fake *FakeCreatedContainer) IsHijackedReturns(result1 bool) {
	fake.IsHijackedStub = nil
	fake.isHijackedReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeCreatedContainer) IsHijackedReturnsOnCall(i int, result1 bool) {
	fake.IsHijackedStub = nil
	if fake.isHijackedReturnsOnCall == nil {
		fake.isHijackedReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isHijackedReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeCreatedContainer) WorkerName() string {
	fake.workerNameMutex.Lock()
	ret, specificReturn := fake.workerNameReturnsOnCall[len(fake.workerNameArgsForCall)]
	fake.workerNameArgsForCall = append(fake.workerNameArgsForCall, struct{}{})
	fake.recordInvocation("WorkerName", []interface{}{})
	fake.workerNameMutex.Unlock()
	if fake.WorkerNameStub != nil {
		return fake.WorkerNameStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.workerNameReturns.result1
}

func (fake *FakeCreatedContainer) WorkerNameCallCount() int {
	fake.workerNameMutex.RLock()
	defer fake.workerNameMutex.RUnlock()
	return len(fake.workerNameArgsForCall)
}

func (fake *FakeCreatedContainer) WorkerNameReturns(result1 string) {
	fake.WorkerNameStub = nil
	fake.workerNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreatedContainer) WorkerNameReturnsOnCall(i int, result1 string) {
	fake.WorkerNameStub = nil
	if fake.workerNameReturnsOnCall == nil {
		fake.workerNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.workerNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreatedContainer) MarkAsHijacked() error {
	fake.markAsHijackedMutex.Lock()
	ret, specificReturn := fake.markAsHijackedReturnsOnCall[len(fake.markAsHijackedArgsForCall)]
	fake.markAsHijackedArgsForCall = append(fake.markAsHijackedArgsForCall, struct{}{})
	fake.recordInvocation("MarkAsHijacked", []interface{}{})
	fake.markAsHijackedMutex.Unlock()
	if fake.MarkAsHijackedStub != nil {
		return fake.MarkAsHijackedStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.markAsHijackedReturns.result1
}

func (fake *FakeCreatedContainer) MarkAsHijackedCallCount() int {
	fake.markAsHijackedMutex.RLock()
	defer fake.markAsHijackedMutex.RUnlock()
	return len(fake.markAsHijackedArgsForCall)
}

func (fake *FakeCreatedContainer) MarkAsHijackedReturns(result1 error) {
	fake.MarkAsHijackedStub = nil
	fake.markAsHijackedReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCreatedContainer) MarkAsHijackedReturnsOnCall(i int, result1 error) {
	fake.MarkAsHijackedStub = nil
	if fake.markAsHijackedReturnsOnCall == nil {
		fake.markAsHijackedReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.markAsHijackedReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCreatedContainer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	fake.discontinueMutex.RLock()
	defer fake.discontinueMutex.RUnlock()
	fake.destroyingMutex.RLock()
	defer fake.destroyingMutex.RUnlock()
	fake.isHijackedMutex.RLock()
	defer fake.isHijackedMutex.RUnlock()
	fake.workerNameMutex.RLock()
	defer fake.workerNameMutex.RUnlock()
	fake.markAsHijackedMutex.RLock()
	defer fake.markAsHijackedMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeCreatedContainer) recordInvocation(key string, args []interface{}) {
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

var _ dbng.CreatedContainer = new(FakeCreatedContainer)