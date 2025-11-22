package telemetry

import (
	"fmt"
	"os"

	"github.com/posthog/posthog-go"
)

const posthogAPIKey = "phc_jkxnO9XCDoPjYwWvSVHfhGPQdn6QIunHRqPsWwZr5Vc" //nolint:gosec // public telemetry API key

// StartTelemetry initializes telemetry with the given version information.
func StartTelemetry(version, commit, buildTime string) {
	client, err := posthog.NewWithConfig(
		posthogAPIKey,
		posthog.Config{Endpoint: "https://eu.i.posthog.com"},
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize telemetry: %v\n", err)

		return
	}
	defer func() {
		if err := client.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "failed to close telemetry client: %v\n", err)
		}
	}()

	if err := client.Enqueue(posthog.Capture{
		DistinctId: "test-user",
		Event:      "app_started",
		Properties: posthog.NewProperties().
			Set("app_version", version).
			Set("commit_hash", commit).
			Set("build_time", buildTime),
	}); err != nil {
		fmt.Fprintf(os.Stderr, "failed to enqueue telemetry event: %v\n", err)
	}
}
