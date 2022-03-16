/*
Copyright hechain. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package etcdraft_test

import (
	"testing"

	"github.com/hechain20/hechain/common/channelconfig"
	"github.com/hechain20/hechain/msp"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestEtcdraft(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Etcdraft Suite")
}

//go:generate counterfeiter -o mocks/orderer_org.go --fake-name OrdererOrg . channelConfigOrdererOrg
type channelConfigOrdererOrg interface {
	channelconfig.OrdererOrg
}

//go:generate counterfeiter -o mocks/msp.go --fake-name MSP . mspInterface
type mspInterface interface {
	msp.MSP
}
