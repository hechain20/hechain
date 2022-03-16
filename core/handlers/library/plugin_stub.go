//go:build noplugin || !cgo
// +build noplugin !cgo

/*
 Copyright hechain All Rights Reserved.

 SPDX-License-Identifier: Apache-2.0
*/

package library

// loadPlugin loads a pluggable handler
func (r *registry) loadPlugin(pluginPath string, handlerType HandlerType, extraArgs ...string) {
	logger.Panicf("Plugins are not supported on this platform")
}
