package subzelle

import (
	"context"
	"flag"

	"github.com/bazelbuild/bazel-gazelle/config"
	"github.com/bazelbuild/bazel-gazelle/rule"

	lpb "github.com/stackb/subzelle/language"
)

// RegisterFlags implements part of the Configurer interface
func (b *subzelle) RegisterFlags(fs *flag.FlagSet, cmd string, c *config.Config) {
	configFlagSet := configFlagSetToProto(c, fs)

	response, err := b.configurerClient.RegisterFlags(context.Background(), &lpb.RegisterFlagsRequest{
		Cmd:           cmd,
		ConfigFlagSet: configFlagSet,
	})
	if err != nil {
		fatalError(err)
	}

	flagSetFromProto(fs, response.Flag)
	configFromProto(c, response.Config)
}

// CheckFlags implements part of the Configurer interface
func (b *subzelle) CheckFlags(fs *flag.FlagSet, c *config.Config) error {
	configFlagSet := configFlagSetToProto(c, fs)

	response, err := b.configurerClient.CheckFlags(context.Background(), configFlagSet)
	if err != nil {
		return err
	}

	flagSetFromProto(fs, response.Flag)
	configFromProto(c, response.Config)

	return nil
}

// KnownDirectives implements part of the Configurer interface
func (b *subzelle) KnownDirectives() []string {
	response, err := b.configurerClient.KnownDirectives(context.Background(), &lpb.KnownDirectivesRequest{})
	if err != nil {
		fatalError(err)
	}

	return response.Directive
}

// Configure implements part of the Configurer interface
func (b *subzelle) Configure(c *config.Config, rel string, f *rule.File) {
	_, err := b.configurerClient.Configure(context.Background(), &lpb.ConfigureRequest{
		Config: configToProto(c),
		Rel:    rel,
		File:   fileToProto(f),
	})
	if err != nil {
		fatalError(err)
	}
}
