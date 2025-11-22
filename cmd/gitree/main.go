package main

import (
	"fmt"
	"os"

	"github.com/andreygrechin/gitree/internal/telemetry"
)

//nolint:gochecknoglobals // Version info variables
var (
	// Version information (injected at build time via ldflags).
	version   = "dev"
	commit    = "none"
	buildTime = "unknown"
)

func main() {
	telemetry.StartTelemetry(version, commit, buildTime)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
