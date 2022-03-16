// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/hechain20/hechain/discovery/support/gossip"
	"github.com/hechain20/hechain/gossip/api"
	"github.com/hechain20/hechain/gossip/common"
	"github.com/hechain20/hechain/gossip/discovery"
	"github.com/hechain20/hechain/gossip/protoext"
)

type Gossip struct {
	IdentityInfoStub        func() api.PeerIdentitySet
	identityInfoMutex       sync.RWMutex
	identityInfoArgsForCall []struct {
	}
	identityInfoReturns struct {
		result1 api.PeerIdentitySet
	}
	identityInfoReturnsOnCall map[int]struct {
		result1 api.PeerIdentitySet
	}
	PeersStub        func() []discovery.NetworkMember
	peersMutex       sync.RWMutex
	peersArgsForCall []struct {
	}
	peersReturns struct {
		result1 []discovery.NetworkMember
	}
	peersReturnsOnCall map[int]struct {
		result1 []discovery.NetworkMember
	}
	PeersOfChannelStub        func(common.ChannelID) []discovery.NetworkMember
	peersOfChannelMutex       sync.RWMutex
	peersOfChannelArgsForCall []struct {
		arg1 common.ChannelID
	}
	peersOfChannelReturns struct {
		result1 []discovery.NetworkMember
	}
	peersOfChannelReturnsOnCall map[int]struct {
		result1 []discovery.NetworkMember
	}
	SelfChannelInfoStub        func(common.ChannelID) *protoext.SignedGossipMessage
	selfChannelInfoMutex       sync.RWMutex
	selfChannelInfoArgsForCall []struct {
		arg1 common.ChannelID
	}
	selfChannelInfoReturns struct {
		result1 *protoext.SignedGossipMessage
	}
	selfChannelInfoReturnsOnCall map[int]struct {
		result1 *protoext.SignedGossipMessage
	}
	SelfMembershipInfoStub        func() discovery.NetworkMember
	selfMembershipInfoMutex       sync.RWMutex
	selfMembershipInfoArgsForCall []struct {
	}
	selfMembershipInfoReturns struct {
		result1 discovery.NetworkMember
	}
	selfMembershipInfoReturnsOnCall map[int]struct {
		result1 discovery.NetworkMember
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Gossip) IdentityInfo() api.PeerIdentitySet {
	fake.identityInfoMutex.Lock()
	ret, specificReturn := fake.identityInfoReturnsOnCall[len(fake.identityInfoArgsForCall)]
	fake.identityInfoArgsForCall = append(fake.identityInfoArgsForCall, struct {
	}{})
	fake.recordInvocation("IdentityInfo", []interface{}{})
	fake.identityInfoMutex.Unlock()
	if fake.IdentityInfoStub != nil {
		return fake.IdentityInfoStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.identityInfoReturns
	return fakeReturns.result1
}

func (fake *Gossip) IdentityInfoCallCount() int {
	fake.identityInfoMutex.RLock()
	defer fake.identityInfoMutex.RUnlock()
	return len(fake.identityInfoArgsForCall)
}

func (fake *Gossip) IdentityInfoCalls(stub func() api.PeerIdentitySet) {
	fake.identityInfoMutex.Lock()
	defer fake.identityInfoMutex.Unlock()
	fake.IdentityInfoStub = stub
}

func (fake *Gossip) IdentityInfoReturns(result1 api.PeerIdentitySet) {
	fake.identityInfoMutex.Lock()
	defer fake.identityInfoMutex.Unlock()
	fake.IdentityInfoStub = nil
	fake.identityInfoReturns = struct {
		result1 api.PeerIdentitySet
	}{result1}
}

func (fake *Gossip) IdentityInfoReturnsOnCall(i int, result1 api.PeerIdentitySet) {
	fake.identityInfoMutex.Lock()
	defer fake.identityInfoMutex.Unlock()
	fake.IdentityInfoStub = nil
	if fake.identityInfoReturnsOnCall == nil {
		fake.identityInfoReturnsOnCall = make(map[int]struct {
			result1 api.PeerIdentitySet
		})
	}
	fake.identityInfoReturnsOnCall[i] = struct {
		result1 api.PeerIdentitySet
	}{result1}
}

func (fake *Gossip) Peers() []discovery.NetworkMember {
	fake.peersMutex.Lock()
	ret, specificReturn := fake.peersReturnsOnCall[len(fake.peersArgsForCall)]
	fake.peersArgsForCall = append(fake.peersArgsForCall, struct {
	}{})
	fake.recordInvocation("Peers", []interface{}{})
	fake.peersMutex.Unlock()
	if fake.PeersStub != nil {
		return fake.PeersStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.peersReturns
	return fakeReturns.result1
}

func (fake *Gossip) PeersCallCount() int {
	fake.peersMutex.RLock()
	defer fake.peersMutex.RUnlock()
	return len(fake.peersArgsForCall)
}

func (fake *Gossip) PeersCalls(stub func() []discovery.NetworkMember) {
	fake.peersMutex.Lock()
	defer fake.peersMutex.Unlock()
	fake.PeersStub = stub
}

func (fake *Gossip) PeersReturns(result1 []discovery.NetworkMember) {
	fake.peersMutex.Lock()
	defer fake.peersMutex.Unlock()
	fake.PeersStub = nil
	fake.peersReturns = struct {
		result1 []discovery.NetworkMember
	}{result1}
}

func (fake *Gossip) PeersReturnsOnCall(i int, result1 []discovery.NetworkMember) {
	fake.peersMutex.Lock()
	defer fake.peersMutex.Unlock()
	fake.PeersStub = nil
	if fake.peersReturnsOnCall == nil {
		fake.peersReturnsOnCall = make(map[int]struct {
			result1 []discovery.NetworkMember
		})
	}
	fake.peersReturnsOnCall[i] = struct {
		result1 []discovery.NetworkMember
	}{result1}
}

func (fake *Gossip) PeersOfChannel(arg1 common.ChannelID) []discovery.NetworkMember {
	fake.peersOfChannelMutex.Lock()
	ret, specificReturn := fake.peersOfChannelReturnsOnCall[len(fake.peersOfChannelArgsForCall)]
	fake.peersOfChannelArgsForCall = append(fake.peersOfChannelArgsForCall, struct {
		arg1 common.ChannelID
	}{arg1})
	fake.recordInvocation("PeersOfChannel", []interface{}{arg1})
	fake.peersOfChannelMutex.Unlock()
	if fake.PeersOfChannelStub != nil {
		return fake.PeersOfChannelStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.peersOfChannelReturns
	return fakeReturns.result1
}

func (fake *Gossip) PeersOfChannelCallCount() int {
	fake.peersOfChannelMutex.RLock()
	defer fake.peersOfChannelMutex.RUnlock()
	return len(fake.peersOfChannelArgsForCall)
}

func (fake *Gossip) PeersOfChannelCalls(stub func(common.ChannelID) []discovery.NetworkMember) {
	fake.peersOfChannelMutex.Lock()
	defer fake.peersOfChannelMutex.Unlock()
	fake.PeersOfChannelStub = stub
}

func (fake *Gossip) PeersOfChannelArgsForCall(i int) common.ChannelID {
	fake.peersOfChannelMutex.RLock()
	defer fake.peersOfChannelMutex.RUnlock()
	argsForCall := fake.peersOfChannelArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Gossip) PeersOfChannelReturns(result1 []discovery.NetworkMember) {
	fake.peersOfChannelMutex.Lock()
	defer fake.peersOfChannelMutex.Unlock()
	fake.PeersOfChannelStub = nil
	fake.peersOfChannelReturns = struct {
		result1 []discovery.NetworkMember
	}{result1}
}

func (fake *Gossip) PeersOfChannelReturnsOnCall(i int, result1 []discovery.NetworkMember) {
	fake.peersOfChannelMutex.Lock()
	defer fake.peersOfChannelMutex.Unlock()
	fake.PeersOfChannelStub = nil
	if fake.peersOfChannelReturnsOnCall == nil {
		fake.peersOfChannelReturnsOnCall = make(map[int]struct {
			result1 []discovery.NetworkMember
		})
	}
	fake.peersOfChannelReturnsOnCall[i] = struct {
		result1 []discovery.NetworkMember
	}{result1}
}

func (fake *Gossip) SelfChannelInfo(arg1 common.ChannelID) *protoext.SignedGossipMessage {
	fake.selfChannelInfoMutex.Lock()
	ret, specificReturn := fake.selfChannelInfoReturnsOnCall[len(fake.selfChannelInfoArgsForCall)]
	fake.selfChannelInfoArgsForCall = append(fake.selfChannelInfoArgsForCall, struct {
		arg1 common.ChannelID
	}{arg1})
	fake.recordInvocation("SelfChannelInfo", []interface{}{arg1})
	fake.selfChannelInfoMutex.Unlock()
	if fake.SelfChannelInfoStub != nil {
		return fake.SelfChannelInfoStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.selfChannelInfoReturns
	return fakeReturns.result1
}

func (fake *Gossip) SelfChannelInfoCallCount() int {
	fake.selfChannelInfoMutex.RLock()
	defer fake.selfChannelInfoMutex.RUnlock()
	return len(fake.selfChannelInfoArgsForCall)
}

func (fake *Gossip) SelfChannelInfoCalls(stub func(common.ChannelID) *protoext.SignedGossipMessage) {
	fake.selfChannelInfoMutex.Lock()
	defer fake.selfChannelInfoMutex.Unlock()
	fake.SelfChannelInfoStub = stub
}

func (fake *Gossip) SelfChannelInfoArgsForCall(i int) common.ChannelID {
	fake.selfChannelInfoMutex.RLock()
	defer fake.selfChannelInfoMutex.RUnlock()
	argsForCall := fake.selfChannelInfoArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Gossip) SelfChannelInfoReturns(result1 *protoext.SignedGossipMessage) {
	fake.selfChannelInfoMutex.Lock()
	defer fake.selfChannelInfoMutex.Unlock()
	fake.SelfChannelInfoStub = nil
	fake.selfChannelInfoReturns = struct {
		result1 *protoext.SignedGossipMessage
	}{result1}
}

func (fake *Gossip) SelfChannelInfoReturnsOnCall(i int, result1 *protoext.SignedGossipMessage) {
	fake.selfChannelInfoMutex.Lock()
	defer fake.selfChannelInfoMutex.Unlock()
	fake.SelfChannelInfoStub = nil
	if fake.selfChannelInfoReturnsOnCall == nil {
		fake.selfChannelInfoReturnsOnCall = make(map[int]struct {
			result1 *protoext.SignedGossipMessage
		})
	}
	fake.selfChannelInfoReturnsOnCall[i] = struct {
		result1 *protoext.SignedGossipMessage
	}{result1}
}

func (fake *Gossip) SelfMembershipInfo() discovery.NetworkMember {
	fake.selfMembershipInfoMutex.Lock()
	ret, specificReturn := fake.selfMembershipInfoReturnsOnCall[len(fake.selfMembershipInfoArgsForCall)]
	fake.selfMembershipInfoArgsForCall = append(fake.selfMembershipInfoArgsForCall, struct {
	}{})
	fake.recordInvocation("SelfMembershipInfo", []interface{}{})
	fake.selfMembershipInfoMutex.Unlock()
	if fake.SelfMembershipInfoStub != nil {
		return fake.SelfMembershipInfoStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.selfMembershipInfoReturns
	return fakeReturns.result1
}

func (fake *Gossip) SelfMembershipInfoCallCount() int {
	fake.selfMembershipInfoMutex.RLock()
	defer fake.selfMembershipInfoMutex.RUnlock()
	return len(fake.selfMembershipInfoArgsForCall)
}

func (fake *Gossip) SelfMembershipInfoCalls(stub func() discovery.NetworkMember) {
	fake.selfMembershipInfoMutex.Lock()
	defer fake.selfMembershipInfoMutex.Unlock()
	fake.SelfMembershipInfoStub = stub
}

func (fake *Gossip) SelfMembershipInfoReturns(result1 discovery.NetworkMember) {
	fake.selfMembershipInfoMutex.Lock()
	defer fake.selfMembershipInfoMutex.Unlock()
	fake.SelfMembershipInfoStub = nil
	fake.selfMembershipInfoReturns = struct {
		result1 discovery.NetworkMember
	}{result1}
}

func (fake *Gossip) SelfMembershipInfoReturnsOnCall(i int, result1 discovery.NetworkMember) {
	fake.selfMembershipInfoMutex.Lock()
	defer fake.selfMembershipInfoMutex.Unlock()
	fake.SelfMembershipInfoStub = nil
	if fake.selfMembershipInfoReturnsOnCall == nil {
		fake.selfMembershipInfoReturnsOnCall = make(map[int]struct {
			result1 discovery.NetworkMember
		})
	}
	fake.selfMembershipInfoReturnsOnCall[i] = struct {
		result1 discovery.NetworkMember
	}{result1}
}

func (fake *Gossip) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.identityInfoMutex.RLock()
	defer fake.identityInfoMutex.RUnlock()
	fake.peersMutex.RLock()
	defer fake.peersMutex.RUnlock()
	fake.peersOfChannelMutex.RLock()
	defer fake.peersOfChannelMutex.RUnlock()
	fake.selfChannelInfoMutex.RLock()
	defer fake.selfChannelInfoMutex.RUnlock()
	fake.selfMembershipInfoMutex.RLock()
	defer fake.selfMembershipInfoMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Gossip) recordInvocation(key string, args []interface{}) {
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

var _ gossip.Gossip = new(Gossip)