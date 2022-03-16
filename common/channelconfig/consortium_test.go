/*
Copyright hechain. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package channelconfig

import (
	"testing"

	"github.com/hechain20/hechain/bccsp/sw"
	"github.com/hechain20/hechain/msp"
	cb "github.com/hyperledger/fabric-protos-go/common"
	"github.com/stretchr/testify/require"
)

func TestConsortiumConfig(t *testing.T) {
	cryptoProvider, err := sw.NewDefaultSecurityLevelWithKeystore(sw.NewDummyKeyStore())
	require.NoError(t, err)
	cc, err := NewConsortiumConfig(&cb.ConfigGroup{}, NewMSPConfigHandler(msp.MSPv1_0, cryptoProvider))
	require.NoError(t, err)
	orgs := cc.Organizations()
	require.Equal(t, 0, len(orgs))

	policy := cc.ChannelCreationPolicy()
	require.EqualValues(t, cb.Policy_UNKNOWN, policy.Type, "Expected policy type to be UNKNOWN")
}
