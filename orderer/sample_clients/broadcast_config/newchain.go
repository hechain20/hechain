// Copyright hechain. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/hechain20/hechain/internal/configtxgen/encoder"
	"github.com/hechain20/hechain/internal/configtxgen/genesisconfig"
	"github.com/hechain20/hechain/internal/pkg/identity"
	cb "github.com/hyperledger/fabric-protos-go/common"
)

func newChainRequest(
	consensusType,
	creationPolicy,
	newChannelID string,
	signer identity.SignerSerializer,
) *cb.Envelope {
	env, err := encoder.MakeChannelCreationTransaction(
		newChannelID,
		signer,
		genesisconfig.Load(genesisconfig.SampleSingleMSPChannelProfile),
	)
	if err != nil {
		panic(err)
	}
	return env
}
