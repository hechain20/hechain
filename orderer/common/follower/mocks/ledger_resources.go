// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hechain20/hechain/orderer/common/follower"
)

type LedgerResources struct {
	AppendStub        func(*common.Block) error
	appendMutex       sync.RWMutex
	appendArgsForCall []struct {
		arg1 *common.Block
	}
	appendReturns struct {
		result1 error
	}
	appendReturnsOnCall map[int]struct {
		result1 error
	}
	BlockStub        func(uint64) *common.Block
	blockMutex       sync.RWMutex
	blockArgsForCall []struct {
		arg1 uint64
	}
	blockReturns struct {
		result1 *common.Block
	}
	blockReturnsOnCall map[int]struct {
		result1 *common.Block
	}
	ChannelIDStub        func() string
	channelIDMutex       sync.RWMutex
	channelIDArgsForCall []struct {
	}
	channelIDReturns struct {
		result1 string
	}
	channelIDReturnsOnCall map[int]struct {
		result1 string
	}
	HeightStub        func() uint64
	heightMutex       sync.RWMutex
	heightArgsForCall []struct {
	}
	heightReturns struct {
		result1 uint64
	}
	heightReturnsOnCall map[int]struct {
		result1 uint64
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *LedgerResources) Append(arg1 *common.Block) error {
	fake.appendMutex.Lock()
	ret, specificReturn := fake.appendReturnsOnCall[len(fake.appendArgsForCall)]
	fake.appendArgsForCall = append(fake.appendArgsForCall, struct {
		arg1 *common.Block
	}{arg1})
	fake.recordInvocation("Append", []interface{}{arg1})
	fake.appendMutex.Unlock()
	if fake.AppendStub != nil {
		return fake.AppendStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.appendReturns
	return fakeReturns.result1
}

func (fake *LedgerResources) AppendCallCount() int {
	fake.appendMutex.RLock()
	defer fake.appendMutex.RUnlock()
	return len(fake.appendArgsForCall)
}

func (fake *LedgerResources) AppendCalls(stub func(*common.Block) error) {
	fake.appendMutex.Lock()
	defer fake.appendMutex.Unlock()
	fake.AppendStub = stub
}

func (fake *LedgerResources) AppendArgsForCall(i int) *common.Block {
	fake.appendMutex.RLock()
	defer fake.appendMutex.RUnlock()
	argsForCall := fake.appendArgsForCall[i]
	return argsForCall.arg1
}

func (fake *LedgerResources) AppendReturns(result1 error) {
	fake.appendMutex.Lock()
	defer fake.appendMutex.Unlock()
	fake.AppendStub = nil
	fake.appendReturns = struct {
		result1 error
	}{result1}
}

func (fake *LedgerResources) AppendReturnsOnCall(i int, result1 error) {
	fake.appendMutex.Lock()
	defer fake.appendMutex.Unlock()
	fake.AppendStub = nil
	if fake.appendReturnsOnCall == nil {
		fake.appendReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.appendReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *LedgerResources) Block(arg1 uint64) *common.Block {
	fake.blockMutex.Lock()
	ret, specificReturn := fake.blockReturnsOnCall[len(fake.blockArgsForCall)]
	fake.blockArgsForCall = append(fake.blockArgsForCall, struct {
		arg1 uint64
	}{arg1})
	fake.recordInvocation("Block", []interface{}{arg1})
	fake.blockMutex.Unlock()
	if fake.BlockStub != nil {
		return fake.BlockStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.blockReturns
	return fakeReturns.result1
}

func (fake *LedgerResources) BlockCallCount() int {
	fake.blockMutex.RLock()
	defer fake.blockMutex.RUnlock()
	return len(fake.blockArgsForCall)
}

func (fake *LedgerResources) BlockCalls(stub func(uint64) *common.Block) {
	fake.blockMutex.Lock()
	defer fake.blockMutex.Unlock()
	fake.BlockStub = stub
}

func (fake *LedgerResources) BlockArgsForCall(i int) uint64 {
	fake.blockMutex.RLock()
	defer fake.blockMutex.RUnlock()
	argsForCall := fake.blockArgsForCall[i]
	return argsForCall.arg1
}

func (fake *LedgerResources) BlockReturns(result1 *common.Block) {
	fake.blockMutex.Lock()
	defer fake.blockMutex.Unlock()
	fake.BlockStub = nil
	fake.blockReturns = struct {
		result1 *common.Block
	}{result1}
}

func (fake *LedgerResources) BlockReturnsOnCall(i int, result1 *common.Block) {
	fake.blockMutex.Lock()
	defer fake.blockMutex.Unlock()
	fake.BlockStub = nil
	if fake.blockReturnsOnCall == nil {
		fake.blockReturnsOnCall = make(map[int]struct {
			result1 *common.Block
		})
	}
	fake.blockReturnsOnCall[i] = struct {
		result1 *common.Block
	}{result1}
}

func (fake *LedgerResources) ChannelID() string {
	fake.channelIDMutex.Lock()
	ret, specificReturn := fake.channelIDReturnsOnCall[len(fake.channelIDArgsForCall)]
	fake.channelIDArgsForCall = append(fake.channelIDArgsForCall, struct {
	}{})
	fake.recordInvocation("ChannelID", []interface{}{})
	fake.channelIDMutex.Unlock()
	if fake.ChannelIDStub != nil {
		return fake.ChannelIDStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.channelIDReturns
	return fakeReturns.result1
}

func (fake *LedgerResources) ChannelIDCallCount() int {
	fake.channelIDMutex.RLock()
	defer fake.channelIDMutex.RUnlock()
	return len(fake.channelIDArgsForCall)
}

func (fake *LedgerResources) ChannelIDCalls(stub func() string) {
	fake.channelIDMutex.Lock()
	defer fake.channelIDMutex.Unlock()
	fake.ChannelIDStub = stub
}

func (fake *LedgerResources) ChannelIDReturns(result1 string) {
	fake.channelIDMutex.Lock()
	defer fake.channelIDMutex.Unlock()
	fake.ChannelIDStub = nil
	fake.channelIDReturns = struct {
		result1 string
	}{result1}
}

func (fake *LedgerResources) ChannelIDReturnsOnCall(i int, result1 string) {
	fake.channelIDMutex.Lock()
	defer fake.channelIDMutex.Unlock()
	fake.ChannelIDStub = nil
	if fake.channelIDReturnsOnCall == nil {
		fake.channelIDReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.channelIDReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *LedgerResources) Height() uint64 {
	fake.heightMutex.Lock()
	ret, specificReturn := fake.heightReturnsOnCall[len(fake.heightArgsForCall)]
	fake.heightArgsForCall = append(fake.heightArgsForCall, struct {
	}{})
	fake.recordInvocation("Height", []interface{}{})
	fake.heightMutex.Unlock()
	if fake.HeightStub != nil {
		return fake.HeightStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.heightReturns
	return fakeReturns.result1
}

func (fake *LedgerResources) HeightCallCount() int {
	fake.heightMutex.RLock()
	defer fake.heightMutex.RUnlock()
	return len(fake.heightArgsForCall)
}

func (fake *LedgerResources) HeightCalls(stub func() uint64) {
	fake.heightMutex.Lock()
	defer fake.heightMutex.Unlock()
	fake.HeightStub = stub
}

func (fake *LedgerResources) HeightReturns(result1 uint64) {
	fake.heightMutex.Lock()
	defer fake.heightMutex.Unlock()
	fake.HeightStub = nil
	fake.heightReturns = struct {
		result1 uint64
	}{result1}
}

func (fake *LedgerResources) HeightReturnsOnCall(i int, result1 uint64) {
	fake.heightMutex.Lock()
	defer fake.heightMutex.Unlock()
	fake.HeightStub = nil
	if fake.heightReturnsOnCall == nil {
		fake.heightReturnsOnCall = make(map[int]struct {
			result1 uint64
		})
	}
	fake.heightReturnsOnCall[i] = struct {
		result1 uint64
	}{result1}
}

func (fake *LedgerResources) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.appendMutex.RLock()
	defer fake.appendMutex.RUnlock()
	fake.blockMutex.RLock()
	defer fake.blockMutex.RUnlock()
	fake.channelIDMutex.RLock()
	defer fake.channelIDMutex.RUnlock()
	fake.heightMutex.RLock()
	defer fake.heightMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *LedgerResources) recordInvocation(key string, args []interface{}) {
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

var _ follower.LedgerResources = new(LedgerResources)
