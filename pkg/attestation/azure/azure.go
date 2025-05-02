// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

package azure

import (
	"github.com/edgelesssys/go-azguestattestation/maa"
)

var (
	OSBuild  = "UVC"
	OSType   = "Linux"
	OSDistro = "UVC"
	MaaURL   = "https://sharedeus.eus.attest.azure.net"
)

func InitializeDefaultMAAVars() {
	maa.OSBuild = OSBuild
	maa.OSType = OSType
	maa.OSDistro = OSDistro
}

func InitializeOSVars(build, osType, osDistro string) {
	if build != "" {
		OSBuild = build
	}
	if osType != "" {
		OSType = osType
	}
	if osDistro != "" {
		OSDistro = osDistro
	}
}
