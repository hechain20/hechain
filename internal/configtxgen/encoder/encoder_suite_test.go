/*
Copyright hechain. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package encoder_test

import (
	"testing"

	"github.com/hechain20/hechain/bccsp/factory"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestEncoder(t *testing.T) {
	factory.InitFactories(nil)

	RegisterFailHandler(Fail)
	RunSpecs(t, "Encoder Suite")
}
