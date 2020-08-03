package app

import (
	"fmt"
	"runtime"
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
	fmt.Printf(`App %s, Compiler: %s %s, Copyright (C) 2020 Vimicasa.`,
		version,
		runtime.Compiler,
		runtime.Version())
	fmt.Println()
}
