package app

import (
	"fmt"
	"runtime"

	"github.com/vimicasa/go-api-demo-skeleton/config"
)

var version string

// SetVersion for setup version string.
func SetVersion(ver string) {
	version = ver
}

// GetVersion for get current version.
func GetVersion() string {
	return version
}

// PrintVersion provide print server engine
func PrintVersion() {
	fmt.Printf(`%s %s, Compiler: %s %s, Copyright (C) 2020 Vimicasa.`,
		config.NameApp,
		version,
		runtime.Compiler,
		runtime.Version())
	fmt.Println()
}
