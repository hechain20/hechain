//go:build race
// +build race

/*
Copyright hechain. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package library

func init() {
	raceEnabled = true
}
