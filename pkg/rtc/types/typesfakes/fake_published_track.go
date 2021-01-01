// Code generated by counterfeiter. DO NOT EDIT.
package typesfakes

import (
	"sync"

	"github.com/livekit/livekit-server/pkg/rtc/types"
	"github.com/livekit/livekit-server/proto/livekit"
)

type FakePublishedTrack struct {
	AddSubscriberStub        func(types.Participant) error
	addSubscriberMutex       sync.RWMutex
	addSubscriberArgsForCall []struct {
		arg1 types.Participant
	}
	addSubscriberReturns struct {
		result1 error
	}
	addSubscriberReturnsOnCall map[int]struct {
		result1 error
	}
	IDStub        func() string
	iDMutex       sync.RWMutex
	iDArgsForCall []struct {
	}
	iDReturns struct {
		result1 string
	}
	iDReturnsOnCall map[int]struct {
		result1 string
	}
	IsMutedStub        func() bool
	isMutedMutex       sync.RWMutex
	isMutedArgsForCall []struct {
	}
	isMutedReturns struct {
		result1 bool
	}
	isMutedReturnsOnCall map[int]struct {
		result1 bool
	}
	KindStub        func() livekit.TrackInfo_Type
	kindMutex       sync.RWMutex
	kindArgsForCall []struct {
	}
	kindReturns struct {
		result1 livekit.TrackInfo_Type
	}
	kindReturnsOnCall map[int]struct {
		result1 livekit.TrackInfo_Type
	}
	RemoveAllSubscribersStub        func()
	removeAllSubscribersMutex       sync.RWMutex
	removeAllSubscribersArgsForCall []struct {
	}
	RemoveSubscriberStub        func(string)
	removeSubscriberMutex       sync.RWMutex
	removeSubscriberArgsForCall []struct {
		arg1 string
	}
	StartStub        func()
	startMutex       sync.RWMutex
	startArgsForCall []struct {
	}
	StreamIDStub        func() string
	streamIDMutex       sync.RWMutex
	streamIDArgsForCall []struct {
	}
	streamIDReturns struct {
		result1 string
	}
	streamIDReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePublishedTrack) AddSubscriber(arg1 types.Participant) error {
	fake.addSubscriberMutex.Lock()
	ret, specificReturn := fake.addSubscriberReturnsOnCall[len(fake.addSubscriberArgsForCall)]
	fake.addSubscriberArgsForCall = append(fake.addSubscriberArgsForCall, struct {
		arg1 types.Participant
	}{arg1})
	stub := fake.AddSubscriberStub
	fakeReturns := fake.addSubscriberReturns
	fake.recordInvocation("AddSubscriber", []interface{}{arg1})
	fake.addSubscriberMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePublishedTrack) AddSubscriberCallCount() int {
	fake.addSubscriberMutex.RLock()
	defer fake.addSubscriberMutex.RUnlock()
	return len(fake.addSubscriberArgsForCall)
}

func (fake *FakePublishedTrack) AddSubscriberCalls(stub func(types.Participant) error) {
	fake.addSubscriberMutex.Lock()
	defer fake.addSubscriberMutex.Unlock()
	fake.AddSubscriberStub = stub
}

func (fake *FakePublishedTrack) AddSubscriberArgsForCall(i int) types.Participant {
	fake.addSubscriberMutex.RLock()
	defer fake.addSubscriberMutex.RUnlock()
	argsForCall := fake.addSubscriberArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePublishedTrack) AddSubscriberReturns(result1 error) {
	fake.addSubscriberMutex.Lock()
	defer fake.addSubscriberMutex.Unlock()
	fake.AddSubscriberStub = nil
	fake.addSubscriberReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePublishedTrack) AddSubscriberReturnsOnCall(i int, result1 error) {
	fake.addSubscriberMutex.Lock()
	defer fake.addSubscriberMutex.Unlock()
	fake.AddSubscriberStub = nil
	if fake.addSubscriberReturnsOnCall == nil {
		fake.addSubscriberReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addSubscriberReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePublishedTrack) ID() string {
	fake.iDMutex.Lock()
	ret, specificReturn := fake.iDReturnsOnCall[len(fake.iDArgsForCall)]
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct {
	}{})
	stub := fake.IDStub
	fakeReturns := fake.iDReturns
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePublishedTrack) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakePublishedTrack) IDCalls(stub func() string) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = stub
}

func (fake *FakePublishedTrack) IDReturns(result1 string) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePublishedTrack) IDReturnsOnCall(i int, result1 string) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakePublishedTrack) IsMuted() bool {
	fake.isMutedMutex.Lock()
	ret, specificReturn := fake.isMutedReturnsOnCall[len(fake.isMutedArgsForCall)]
	fake.isMutedArgsForCall = append(fake.isMutedArgsForCall, struct {
	}{})
	stub := fake.IsMutedStub
	fakeReturns := fake.isMutedReturns
	fake.recordInvocation("IsMuted", []interface{}{})
	fake.isMutedMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePublishedTrack) IsMutedCallCount() int {
	fake.isMutedMutex.RLock()
	defer fake.isMutedMutex.RUnlock()
	return len(fake.isMutedArgsForCall)
}

func (fake *FakePublishedTrack) IsMutedCalls(stub func() bool) {
	fake.isMutedMutex.Lock()
	defer fake.isMutedMutex.Unlock()
	fake.IsMutedStub = stub
}

func (fake *FakePublishedTrack) IsMutedReturns(result1 bool) {
	fake.isMutedMutex.Lock()
	defer fake.isMutedMutex.Unlock()
	fake.IsMutedStub = nil
	fake.isMutedReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakePublishedTrack) IsMutedReturnsOnCall(i int, result1 bool) {
	fake.isMutedMutex.Lock()
	defer fake.isMutedMutex.Unlock()
	fake.IsMutedStub = nil
	if fake.isMutedReturnsOnCall == nil {
		fake.isMutedReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isMutedReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakePublishedTrack) Kind() livekit.TrackInfo_Type {
	fake.kindMutex.Lock()
	ret, specificReturn := fake.kindReturnsOnCall[len(fake.kindArgsForCall)]
	fake.kindArgsForCall = append(fake.kindArgsForCall, struct {
	}{})
	stub := fake.KindStub
	fakeReturns := fake.kindReturns
	fake.recordInvocation("Kind", []interface{}{})
	fake.kindMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePublishedTrack) KindCallCount() int {
	fake.kindMutex.RLock()
	defer fake.kindMutex.RUnlock()
	return len(fake.kindArgsForCall)
}

func (fake *FakePublishedTrack) KindCalls(stub func() livekit.TrackInfo_Type) {
	fake.kindMutex.Lock()
	defer fake.kindMutex.Unlock()
	fake.KindStub = stub
}

func (fake *FakePublishedTrack) KindReturns(result1 livekit.TrackInfo_Type) {
	fake.kindMutex.Lock()
	defer fake.kindMutex.Unlock()
	fake.KindStub = nil
	fake.kindReturns = struct {
		result1 livekit.TrackInfo_Type
	}{result1}
}

func (fake *FakePublishedTrack) KindReturnsOnCall(i int, result1 livekit.TrackInfo_Type) {
	fake.kindMutex.Lock()
	defer fake.kindMutex.Unlock()
	fake.KindStub = nil
	if fake.kindReturnsOnCall == nil {
		fake.kindReturnsOnCall = make(map[int]struct {
			result1 livekit.TrackInfo_Type
		})
	}
	fake.kindReturnsOnCall[i] = struct {
		result1 livekit.TrackInfo_Type
	}{result1}
}

func (fake *FakePublishedTrack) RemoveAllSubscribers() {
	fake.removeAllSubscribersMutex.Lock()
	fake.removeAllSubscribersArgsForCall = append(fake.removeAllSubscribersArgsForCall, struct {
	}{})
	stub := fake.RemoveAllSubscribersStub
	fake.recordInvocation("RemoveAllSubscribers", []interface{}{})
	fake.removeAllSubscribersMutex.Unlock()
	if stub != nil {
		fake.RemoveAllSubscribersStub()
	}
}

func (fake *FakePublishedTrack) RemoveAllSubscribersCallCount() int {
	fake.removeAllSubscribersMutex.RLock()
	defer fake.removeAllSubscribersMutex.RUnlock()
	return len(fake.removeAllSubscribersArgsForCall)
}

func (fake *FakePublishedTrack) RemoveAllSubscribersCalls(stub func()) {
	fake.removeAllSubscribersMutex.Lock()
	defer fake.removeAllSubscribersMutex.Unlock()
	fake.RemoveAllSubscribersStub = stub
}

func (fake *FakePublishedTrack) RemoveSubscriber(arg1 string) {
	fake.removeSubscriberMutex.Lock()
	fake.removeSubscriberArgsForCall = append(fake.removeSubscriberArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.RemoveSubscriberStub
	fake.recordInvocation("RemoveSubscriber", []interface{}{arg1})
	fake.removeSubscriberMutex.Unlock()
	if stub != nil {
		fake.RemoveSubscriberStub(arg1)
	}
}

func (fake *FakePublishedTrack) RemoveSubscriberCallCount() int {
	fake.removeSubscriberMutex.RLock()
	defer fake.removeSubscriberMutex.RUnlock()
	return len(fake.removeSubscriberArgsForCall)
}

func (fake *FakePublishedTrack) RemoveSubscriberCalls(stub func(string)) {
	fake.removeSubscriberMutex.Lock()
	defer fake.removeSubscriberMutex.Unlock()
	fake.RemoveSubscriberStub = stub
}

func (fake *FakePublishedTrack) RemoveSubscriberArgsForCall(i int) string {
	fake.removeSubscriberMutex.RLock()
	defer fake.removeSubscriberMutex.RUnlock()
	argsForCall := fake.removeSubscriberArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePublishedTrack) Start() {
	fake.startMutex.Lock()
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
	}{})
	stub := fake.StartStub
	fake.recordInvocation("Start", []interface{}{})
	fake.startMutex.Unlock()
	if stub != nil {
		fake.StartStub()
	}
}

func (fake *FakePublishedTrack) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakePublishedTrack) StartCalls(stub func()) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = stub
}

func (fake *FakePublishedTrack) StreamID() string {
	fake.streamIDMutex.Lock()
	ret, specificReturn := fake.streamIDReturnsOnCall[len(fake.streamIDArgsForCall)]
	fake.streamIDArgsForCall = append(fake.streamIDArgsForCall, struct {
	}{})
	stub := fake.StreamIDStub
	fakeReturns := fake.streamIDReturns
	fake.recordInvocation("StreamID", []interface{}{})
	fake.streamIDMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePublishedTrack) StreamIDCallCount() int {
	fake.streamIDMutex.RLock()
	defer fake.streamIDMutex.RUnlock()
	return len(fake.streamIDArgsForCall)
}

func (fake *FakePublishedTrack) StreamIDCalls(stub func() string) {
	fake.streamIDMutex.Lock()
	defer fake.streamIDMutex.Unlock()
	fake.StreamIDStub = stub
}

func (fake *FakePublishedTrack) StreamIDReturns(result1 string) {
	fake.streamIDMutex.Lock()
	defer fake.streamIDMutex.Unlock()
	fake.StreamIDStub = nil
	fake.streamIDReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePublishedTrack) StreamIDReturnsOnCall(i int, result1 string) {
	fake.streamIDMutex.Lock()
	defer fake.streamIDMutex.Unlock()
	fake.StreamIDStub = nil
	if fake.streamIDReturnsOnCall == nil {
		fake.streamIDReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.streamIDReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakePublishedTrack) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addSubscriberMutex.RLock()
	defer fake.addSubscriberMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.isMutedMutex.RLock()
	defer fake.isMutedMutex.RUnlock()
	fake.kindMutex.RLock()
	defer fake.kindMutex.RUnlock()
	fake.removeAllSubscribersMutex.RLock()
	defer fake.removeAllSubscribersMutex.RUnlock()
	fake.removeSubscriberMutex.RLock()
	defer fake.removeSubscriberMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.streamIDMutex.RLock()
	defer fake.streamIDMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePublishedTrack) recordInvocation(key string, args []interface{}) {
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

var _ types.PublishedTrack = new(FakePublishedTrack)