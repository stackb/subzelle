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
	plugin           *PluginConfig
	languageClient   lpb.LanguageClient
	configurerClient lpb.ConfigurerClient
	conn             *grpc.ClientConn
	cmd              *exec.Cmd
}

// NewLanguage is the entrypoint for gazelle language plugin
func NewLanguage() language.Language {
	var cmd *exec.Cmd
	var err error

	plugin := GetPluginConfig()
	address := plugin.Address
	if address == "" {
		address = "localhost:50051"
		log.Printf("Launching subprocess: %s (%s)", plugin.Executable, address)

		cmd, err = startPlugin(".", plugin.Executable, nil, []string{
			fmt.Sprintf("%sADDRESS=%s", PluginEnvVarNamePrefix, address),
		})
		if err != nil {
			fatalError(fmt.Errorf("could not start plugin %q: %v", plugin.Executable, err))
		}
	} else {
		log.Fatalf("connection address: %v", address)
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fatalError(err)
	}

	return &subzelle{
		plugin:           plugin,
		languageClient:   lpb.NewLanguageClient(conn),
		configurerClient: lpb.NewConfigurerClient(conn),
		conn:             conn,
		cmd:              cmd,
	}
}
