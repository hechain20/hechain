// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hechain20/hechain/orderer/common/channelparticipation"
	"github.com/hechain20/hechain/orderer/common/types"
)

type ChannelManagement struct {
	ChannelInfoStub        func(string) (types.ChannelInfo, error)
	channelInfoMutex       sync.RWMutex
	channelInfoArgsForCall []struct {
		arg1 string
	}
	channelInfoReturns struct {
		result1 types.ChannelInfo
		result2 error
	}
	channelInfoReturnsOnCall map[int]struct {
		result1 types.ChannelInfo
		result2 error
	}
	ChannelListStub        func() types.ChannelList
	channelListMutex       sync.RWMutex
	channelListArgsForCall []struct {
	}
	channelListReturns struct {
		result1 types.ChannelList
	}
	channelListReturnsOnCall map[int]struct {
		result1 types.ChannelList
	}
	JoinChannelStub        func(string, *common.Block, bool) (types.ChannelInfo, error)
	joinChannelMutex       sync.RWMutex
	joinChannelArgsForCall []struct {
		arg1 string
		arg2 *common.Block
		arg3 bool
	}
	joinChannelReturns struct {
		result1 types.ChannelInfo
		result2 error
	}
	joinChannelReturnsOnCall map[int]struct {
		result1 types.ChannelInfo
		result2 error
	}
	RemoveChannelStub        func(string) error
	removeChannelMutex       sync.RWMutex
	removeChannelArgsForCall []struct {
		arg1 string
	}
	removeChannelReturns struct {
		result1 error
	}
	removeChannelReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ChannelManagement) ChannelInfo(arg1 string) (types.ChannelInfo, error) {
	fake.channelInfoMutex.Lock()
	ret, specificReturn := fake.channelInfoReturnsOnCall[len(fake.channelInfoArgsForCall)]
	fake.channelInfoArgsForCall = append(fake.channelInfoArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ChannelInfo", []interface{}{arg1})
	fake.channelInfoMutex.Unlock()
	if fake.ChannelInfoStub != nil {
		return fake.ChannelInfoStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.channelInfoReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ChannelManagement) ChannelInfoCallCount() int {
	fake.channelInfoMutex.RLock()
	defer fake.channelInfoMutex.RUnlock()
	return len(fake.channelInfoArgsForCall)
}

func (fake *ChannelManagement) ChannelInfoCalls(stub func(string) (types.ChannelInfo, error)) {
	fake.channelInfoMutex.Lock()
	defer fake.channelInfoMutex.Unlock()
	fake.ChannelInfoStub = stub
}

func (fake *ChannelManagement) ChannelInfoArgsForCall(i int) string {
	fake.channelInfoMutex.RLock()
	defer fake.channelInfoMutex.RUnlock()
	argsForCall := fake.channelInfoArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ChannelManagement) ChannelInfoReturns(result1 types.ChannelInfo, result2 error) {
	fake.channelInfoMutex.Lock()
	defer fake.channelInfoMutex.Unlock()
	fake.ChannelInfoStub = nil
	fake.channelInfoReturns = struct {
		result1 types.ChannelInfo
		result2 error
	}{result1, result2}
}

func (fake *ChannelManagement) ChannelInfoReturnsOnCall(i int, result1 types.ChannelInfo, result2 error) {
	fake.channelInfoMutex.Lock()
	defer fake.channelInfoMutex.Unlock()
	fake.ChannelInfoStub = nil
	if fake.channelInfoReturnsOnCall == nil {
		fake.channelInfoReturnsOnCall = make(map[int]struct {
			result1 types.ChannelInfo
			result2 error
		})
	}
	fake.channelInfoReturnsOnCall[i] = struct {
		result1 types.ChannelInfo
		result2 error
	}{result1, result2}
}

func (fake *ChannelManagement) ChannelList() types.ChannelList {
	fake.channelListMutex.Lock()
	ret, specificReturn := fake.channelListReturnsOnCall[len(fake.channelListArgsForCall)]
	fake.channelListArgsForCall = append(fake.channelListArgsForCall, struct {
	}{})
	fake.recordInvocation("ChannelList", []interface{}{})
	fake.channelListMutex.Unlock()
	if fake.ChannelListStub != nil {
		return fake.ChannelListStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.channelListReturns
	return fakeReturns.result1
}

func (fake *ChannelManagement) ChannelListCallCount() int {
	fake.channelListMutex.RLock()
	defer fake.channelListMutex.RUnlock()
	return len(fake.channelListArgsForCall)
}

func (fake *ChannelManagement) ChannelListCalls(stub func() types.ChannelList) {
	fake.channelListMutex.Lock()
	defer fake.channelListMutex.Unlock()
	fake.ChannelListStub = stub
}

func (fake *ChannelManagement) ChannelListReturns(result1 types.ChannelList) {
	fake.channelListMutex.Lock()
	defer fake.channelListMutex.Unlock()
	fake.ChannelListStub = nil
	fake.channelListReturns = struct {
		result1 types.ChannelList
	}{result1}
}

func (fake *ChannelManagement) ChannelListReturnsOnCall(i int, result1 types.ChannelList) {
	fake.channelListMutex.Lock()
	defer fake.channelListMutex.Unlock()
	fake.ChannelListStub = nil
	if fake.channelListReturnsOnCall == nil {
		fake.channelListReturnsOnCall = make(map[int]struct {
			result1 types.ChannelList
		})
	}
	fake.channelListReturnsOnCall[i] = struct {
		result1 types.ChannelList
	}{result1}
}

func (fake *ChannelManagement) JoinChannel(arg1 string, arg2 *common.Block, arg3 bool) (types.ChannelInfo, error) {
	fake.joinChannelMutex.Lock()
	ret, specificReturn := fake.joinChannelReturnsOnCall[len(fake.joinChannelArgsForCall)]
	fake.joinChannelArgsForCall = append(fake.joinChannelArgsForCall, struct {
		arg1 string
		arg2 *common.Block
		arg3 bool
	}{arg1, arg2, arg3})
	fake.recordInvocation("JoinChannel", []interface{}{arg1, arg2, arg3})
	fake.joinChannelMutex.Unlock()
	if fake.JoinChannelStub != nil {
		return fake.JoinChannelStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.joinChannelReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ChannelManagement) JoinChannelCallCount() int {
	fake.joinChannelMutex.RLock()
	defer fake.joinChannelMutex.RUnlock()
	return len(fake.joinChannelArgsForCall)
}

func (fake *ChannelManagement) JoinChannelCalls(stub func(string, *common.Block, bool) (types.ChannelInfo, error)) {
	fake.joinChannelMutex.Lock()
	defer fake.joinChannelMutex.Unlock()
	fake.JoinChannelStub = stub
}

func (fake *ChannelManagement) JoinChannelArgsForCall(i int) (string, *common.Block, bool) {
	fake.joinChannelMutex.RLock()
	defer fake.joinChannelMutex.RUnlock()
	argsForCall := fake.joinChannelArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *ChannelManagement) JoinChannelReturns(result1 types.ChannelInfo, result2 error) {
	fake.joinChannelMutex.Lock()
	defer fake.joinChannelMutex.Unlock()
	fake.JoinChannelStub = nil
	fake.joinChannelReturns = struct {
		result1 types.ChannelInfo
		result2 error
	}{result1, result2}
}

func (fake *ChannelManagement) JoinChannelReturnsOnCall(i int, result1 types.ChannelInfo, result2 error) {
	fake.joinChannelMutex.Lock()
	defer fake.joinChannelMutex.Unlock()
	fake.JoinChannelStub = nil
	if fake.joinChannelReturnsOnCall == nil {
		fake.joinChannelReturnsOnCall = make(map[int]struct {
			result1 types.ChannelInfo
			result2 error
		})
	}
	fake.joinChannelReturnsOnCall[i] = struct {
		result1 types.ChannelInfo
		result2 error
	}{result1, result2}
}

func (fake *ChannelManagement) RemoveChannel(arg1 string) error {
	fake.removeChannelMutex.Lock()
	ret, specificReturn := fake.removeChannelReturnsOnCall[len(fake.removeChannelArgsForCall)]
	fake.removeChannelArgsForCall = append(fake.removeChannelArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("RemoveChannel", []interface{}{arg1})
	fake.removeChannelMutex.Unlock()
	if fake.RemoveChannelStub != nil {
		return fake.RemoveChannelStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.removeChannelReturns
	return fakeReturns.result1
}

func (fake *ChannelManagement) RemoveChannelCallCount() int {
	fake.removeChannelMutex.RLock()
	defer fake.removeChannelMutex.RUnlock()
	return len(fake.removeChannelArgsForCall)
}

func (fake *ChannelManagement) RemoveChannelCalls(stub func(string) error) {
	fake.removeChannelMutex.Lock()
	defer fake.removeChannelMutex.Unlock()
	fake.RemoveChannelStub = stub
}

func (fake *ChannelManagement) RemoveChannelArgsForCall(i int) string {
	fake.removeChannelMutex.RLock()
	defer fake.removeChannelMutex.RUnlock()
	argsForCall := fake.removeChannelArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ChannelManagement) RemoveChannelReturns(result1 error) {
	fake.removeChannelMutex.Lock()
	defer fake.removeChannelMutex.Unlock()
	fake.RemoveChannelStub = nil
	fake.removeChannelReturns = struct {
		result1 error
	}{result1}
}

func (fake *ChannelManagement) RemoveChannelReturnsOnCall(i int, result1 error) {
	fake.removeChannelMutex.Lock()
	defer fake.removeChannelMutex.Unlock()
	fake.RemoveChannelStub = nil
	if fake.removeChannelReturnsOnCall == nil {
		fake.removeChannelReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeChannelReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ChannelManagement) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.channelInfoMutex.RLock()
	defer fake.channelInfoMutex.RUnlock()
	fake.channelListMutex.RLock()
	defer fake.channelListMutex.RUnlock()
	fake.joinChannelMutex.RLock()
	defer fake.joinChannelMutex.RUnlock()
	fake.removeChannelMutex.RLock()
	defer fake.removeChannelMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ChannelManagement) recordInvocation(key string, args []interface{}) {
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

var _ channelparticipation.ChannelManagement = new(ChannelManagement)
