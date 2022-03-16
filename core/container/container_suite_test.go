/*
Copyright hechain. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package container_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestContainer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Container Suite")
}
