// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	cluster "github.com/hechain20/hechain/orderer/common/cluster"
	mock "github.com/stretchr/testify/mock"
)

// VerifierRetriever is an autogenerated mock type for the VerifierRetriever type
type VerifierRetriever struct {
	mock.Mock
}

// RetrieveVerifier provides a mock function with given fields: channel
func (_m *VerifierRetriever) RetrieveVerifier(channel string) cluster.BlockVerifier {
	ret := _m.Called(channel)

	var r0 cluster.BlockVerifier
	if rf, ok := ret.Get(0).(func(string) cluster.BlockVerifier); ok {
		r0 = rf(channel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cluster.BlockVerifier)
		}
	}

	return r0
}
