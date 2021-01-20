package bazelle

import (
	"context"
	"flag"
	"fmt"
	"os"

	"google.golang.org/grpc"

	"github.com/bazelbuild/bazel-gazelle/config"
	"github.com/bazelbuild/bazel-gazelle/label"
	"github.com/bazelbuild/bazel-gazelle/language"
	"github.com/bazelbuild/bazel-gazelle/repo"
	"github.com/bazelbuild/bazel-gazelle/resolve"
	"github.com/bazelbuild/bazel-gazelle/rule"

	pb "github.com/stackb/bazelle/proto"
)

type rpcLang struct {
	client pb.LanguageClient
	conn   *grpc.ClientConn
}

// NewLanguage is the entrypoint for gazelle language plugin
func NewLanguage() language.Language {
	cfg := GetActiveConfiguration()
	server := cfg.GetServer()
	if server == nil {
		fatalError(fmt.Errorf("configuration file block 'Server' is mandatory"))
	}
	conn, err := grpc.Dial(server.GetServerAddress(), grpc.WithInsecure())
	if err != nil {
		fatalError(err)
	}

	client := pb.NewLanguageClient(conn)
	return &rpcLang{
		client: client,
		conn:   conn,
	}
}

func (x *rpcLang) Name() string {
	return GetActiveConfiguration().Server.GetLanguageName()
}

func (x *rpcLang) Kinds() map[string]rule.KindInfo {
	return map[string]rule.KindInfo{
		"x_library": {},
	}
}

func (x *rpcLang) Loads() []rule.LoadInfo {
	resp, err := x.client.Loads(context.Background(), &pb.LoadsRequest{})
	if err != nil {
		fatalError(err)
	}
	return makeLoadInfos(resp.Load)
}

// RegisterFlags implements part of the Configurer interface
func (x *rpcLang) RegisterFlags(fs *flag.FlagSet, cmd string, c *config.Config) {
}

// CheckFlags implements part of the Configurer interface
func (x *rpcLang) CheckFlags(fs *flag.FlagSet, c *config.Config) error {
	return nil
}

// KnownDirectives implements part of the Configurer interface
func (x *rpcLang) KnownDirectives() []string {
	return nil
}

func (x *rpcLang) Configure(c *config.Config, rel string, f *rule.File) {
}

func (x *rpcLang) GenerateRules(args language.GenerateArgs) language.GenerateResult {
	return language.GenerateResult{
		Gen:     []*rule.Rule{rule.NewRule("x_library", "x_default_library")},
		Imports: []interface{}{nil},
	}
}

func (x *rpcLang) Fix(c *config.Config, f *rule.File) {
}

func (x *rpcLang) Imports(c *config.Config, r *rule.Rule, f *rule.File) []resolve.ImportSpec {
	return nil
}

func (x *rpcLang) Embeds(r *rule.Rule, from label.Label) []label.Label {
	return nil
}

func (x *rpcLang) Resolve(c *config.Config, ix *resolve.RuleIndex, rc *repo.RemoteCache, r *rule.Rule, imports interface{}, from label.Label) {
}

func makeLoadInfos(loads []*pb.LoadInfo) []rule.LoadInfo {
	infos := make([]rule.LoadInfo, len(loads))
	for i, load := range loads {
		infos[i] = makeLoadInfo(load)
	}
	return infos
}

func makeLoadInfo(load *pb.LoadInfo) rule.LoadInfo {
	return rule.LoadInfo{load.Name, load.Symbols, load.After}
}

func fatalError(err error) {
	fmt.Fprintf(os.Stderr, "could not dial remote: %v", err)
	os.Exit(1)
}
