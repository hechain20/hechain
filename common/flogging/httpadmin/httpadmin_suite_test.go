/*
Copyright hechain. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package httpadmin_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHttpadmin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Httpadmin Suite")
}
