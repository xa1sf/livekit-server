// Code generated by counterfeiter. DO NOT EDIT.
package typesfakes

import (
	"sync"

	"github.com/livekit/livekit-server/pkg/rtc/types"
	"github.com/pion/rtp"
)

type FakePacketBuffer struct {
	GetBufferedPacketsStub        func(uint32, uint16, uint32, []uint16) []rtp.Packet
	getBufferedPacketsMutex       sync.RWMutex
	getBufferedPacketsArgsForCall []struct {
		arg1 uint32
		arg2 uint16
		arg3 uint32
		arg4 []uint16
	}
	getBufferedPacketsReturns struct {
		result1 []rtp.Packet
	}
	getBufferedPacketsReturnsOnCall map[int]struct {
		result1 []rtp.Packet
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePacketBuffer) GetBufferedPackets(arg1 uint32, arg2 uint16, arg3 uint32, arg4 []uint16) []rtp.Packet {
	var arg4Copy []uint16
	if arg4 != nil {
		arg4Copy = make([]uint16, len(arg4))
		copy(arg4Copy, arg4)
	}
	fake.getBufferedPacketsMutex.Lock()
	ret, specificReturn := fake.getBufferedPacketsReturnsOnCall[len(fake.getBufferedPacketsArgsForCall)]
	fake.getBufferedPacketsArgsForCall = append(fake.getBufferedPacketsArgsForCall, struct {
		arg1 uint32
		arg2 uint16
		arg3 uint32
		arg4 []uint16
	}{arg1, arg2, arg3, arg4Copy})
	stub := fake.GetBufferedPacketsStub
	fakeReturns := fake.getBufferedPacketsReturns
	fake.recordInvocation("GetBufferedPackets", []interface{}{arg1, arg2, arg3, arg4Copy})
	fake.getBufferedPacketsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePacketBuffer) GetBufferedPacketsCallCount() int {
	fake.getBufferedPacketsMutex.RLock()
	defer fake.getBufferedPacketsMutex.RUnlock()
	return len(fake.getBufferedPacketsArgsForCall)
}

func (fake *FakePacketBuffer) GetBufferedPacketsCalls(stub func(uint32, uint16, uint32, []uint16) []rtp.Packet) {
	fake.getBufferedPacketsMutex.Lock()
	defer fake.getBufferedPacketsMutex.Unlock()
	fake.GetBufferedPacketsStub = stub
}

func (fake *FakePacketBuffer) GetBufferedPacketsArgsForCall(i int) (uint32, uint16, uint32, []uint16) {
	fake.getBufferedPacketsMutex.RLock()
	defer fake.getBufferedPacketsMutex.RUnlock()
	argsForCall := fake.getBufferedPacketsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakePacketBuffer) GetBufferedPacketsReturns(result1 []rtp.Packet) {
	fake.getBufferedPacketsMutex.Lock()
	defer fake.getBufferedPacketsMutex.Unlock()
	fake.GetBufferedPacketsStub = nil
	fake.getBufferedPacketsReturns = struct {
		result1 []rtp.Packet
	}{result1}
}

func (fake *FakePacketBuffer) GetBufferedPacketsReturnsOnCall(i int, result1 []rtp.Packet) {
	fake.getBufferedPacketsMutex.Lock()
	defer fake.getBufferedPacketsMutex.Unlock()
	fake.GetBufferedPacketsStub = nil
	if fake.getBufferedPacketsReturnsOnCall == nil {
		fake.getBufferedPacketsReturnsOnCall = make(map[int]struct {
			result1 []rtp.Packet
		})
	}
	fake.getBufferedPacketsReturnsOnCall[i] = struct {
		result1 []rtp.Packet
	}{result1}
}

func (fake *FakePacketBuffer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getBufferedPacketsMutex.RLock()
	defer fake.getBufferedPacketsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePacketBuffer) recordInvocation(key string, args []interface{}) {
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

var _ types.PacketBuffer = new(FakePacketBuffer)