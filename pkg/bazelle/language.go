package subzelle

import (
	"google.golang.org/grpc"

	"github.com/bazelbuild/bazel-gazelle/language"

	lpb "github.com/stackb/subzelle/language"
)

type subzelle struct {
	plugin           *PluginConfig
	languageClient   lpb.LanguageClient
	configurerClient lpb.ConfigurerClient
	conn             *grpc.ClientConn
}

// NewLanguage is the entrypoint for gazelle language plugin
func NewLanguage() language.Language {

	plugin := GetPluginConfig()

	conn, err := grpc.Dial(plugin.Address, grpc.WithInsecure())
	if err != nil {
		fatalError(err)
	}

	return &subzelle{
		plugin:           plugin,
		languageClient:   lpb.NewLanguageClient(conn),
		configurerClient: lpb.NewConfigurerClient(conn),
		conn:             conn,
	}
}
