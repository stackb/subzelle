package subzelle

import (
	"fmt"
	"log"
	"os/exec"

	"google.golang.org/grpc"

	"github.com/bazelbuild/bazel-gazelle/language"

	lpb "github.com/stackb/subzelle/language"
)

type subzelle struct {
	plugin *PluginConfig
	client lpb.LanguageClient
	conn   *grpc.ClientConn
	cmd    *exec.Cmd
}

// NewLanguage is the entrypoint for gazelle language plugin
func NewLanguage() language.Language {
	var cmd *exec.Cmd
	var err error

	plugin := GetPluginConfig()
	address := plugin.Address
	if address == "" {
		address = "0.0.0.0:50051"
		log.Printf("Launching subprocess: %s (%s)", plugin.Path, address)
		cmd, err = startPlugin(plugin.Root, plugin.Path, nil, []string{
			fmt.Sprintf("%sADDRESS=%s", PluginEnvVarNamePrefix, address),
		})
		if err != nil {
			fatalError(fmt.Errorf("could not start plugin %q: %v", plugin.Path, err))
		}
	}

	conn, err := grpc.Dial(address,
		grpc.WithInsecure(),
	)

	if err != nil {
		fatalError(err)
	}

	return &subzelle{
		plugin: plugin,
		client: lpb.NewLanguageClient(conn),
		conn:   conn,
		cmd:    cmd,
	}
}
