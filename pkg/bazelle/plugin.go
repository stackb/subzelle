package subzelle

import (
	"fmt"
	"os"
)

var (
	PluginDefaultAddress   = "localhost:50051"
	PluginEnvVarNamePrefix = "subzelle_PLUGIN_"
)

// PluginConfig captures the configuration required to launch and connect to a
// gRPC based plugin.
type PluginConfig struct {
	// The name of the plugin.  This becomes the gazelle language name.
	Name string
	// The binary filename.  If this is not set a plugin subprocess will not be
	// launched.
	Executable string
	// The network address where the plugin is running.  If this is set gazelle
	// will attempt to connect directly to the service without launching a
	// subprocess.
	Address string
}

// GetPluginConfig constructs the plugin configuration from environment
// variables.
func GetPluginConfig() *PluginConfig {
	name := mustGetEnvVar(PluginEnvVarNamePrefix + "NAME")
	executable := maybeGetEnvVar(PluginEnvVarNamePrefix+"EXECUTABLE", "")
	address := maybeGetEnvVar(PluginEnvVarNamePrefix+"ADDRESS", "")

	return &PluginConfig{name, address, executable}
}

func mustGetEnvVar(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		fatalError(fmt.Errorf("Required environment variable %q is not set\n", key))
	}
	return value
}

func maybeGetEnvVar(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return value
}

func fatalError(err error) {
	fmt.Fprintf(os.Stderr, "could not dial remote: %v", err)
	os.Exit(1)
}
