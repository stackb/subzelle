package bazelle

import (
	"flag"

	"github.com/bazelbuild/bazel-gazelle/config"
	"github.com/bazelbuild/bazel-gazelle/language"
	"github.com/bazelbuild/bazel-gazelle/rule"

	lpb "github.com/stackb/bazelle/language"
)

//
// type conversion utility functions
//

func loadInfosFromProto(loads []*lpb.LoadInfo) []rule.LoadInfo {
	infos := make([]rule.LoadInfo, len(loads))
	for i, load := range loads {
		infos[i] = loadInfoFromProto(load)
	}
	return infos
}

func loadInfoFromProto(load *lpb.LoadInfo) rule.LoadInfo {
	return rule.LoadInfo{load.Name, load.Symbols, load.After}
}

func kindInfoFromProto(kind *lpb.KindInfo) rule.KindInfo {
	return rule.KindInfo{
		MatchAny:        kind.MatchAny,
		MatchAttrs:      kind.MatchAttrs,
		NonEmptyAttrs:   kind.NonEmptyAttrs,
		SubstituteAttrs: kind.SubstituteAttrs,
		MergeableAttrs:  kind.MergeableAttrs,
		ResolveAttrs:    kind.ResolveAttrs,
	}
}

func generateArgsToProto(args language.GenerateArgs) *lpb.GenerateArgs {
	return &lpb.GenerateArgs{
		Config:       configToProto(args.Config),
		Dir:          args.Dir,
		Rel:          args.Rel,
		File:         fileToProto(args.File),
		Subdirs:      args.Subdirs,
		RegularFiles: args.RegularFiles,
		GenFiles:     args.GenFiles,
		OtherEmpty:   rulesToProto(args.OtherEmpty),
		OtherGen:     rulesToProto(args.OtherGen),
	}
}

func generateResultFromProto(args *lpb.GenerateResult) language.GenerateResult {
	return language.GenerateResult{
		Gen:   rulesFromProto(args.Gen),
		Empty: rulesFromProto(args.Empty),
	}
}

func configToProto(config *config.Config) *lpb.Config {
	return &lpb.Config{
		RepoRoot:            config.RepoRoot,
		RepoName:            config.RepoName,
		ReadBuildFilesDir:   config.ReadBuildFilesDir,
		WriteBuildFilesDir:  config.WriteBuildFilesDir,
		ValidBuildFileNames: config.ValidBuildFileNames,
		ShouldFix:           config.ShouldFix,
		IndexLibraries:      config.IndexLibraries,
		KindMap:             mappedKindsToProto(config.KindMap),
		Langs:               config.Langs,
	}
}

func mappedKindsToProto(kindMap map[string]config.MappedKind) map[string]*lpb.MappedKind {
	result := make(map[string]*lpb.MappedKind)
	for k, v := range kindMap {
		result[k] = mappedKindToProto(v)
	}
	return result
}

func mappedKindToProto(in config.MappedKind) *lpb.MappedKind {
	return &lpb.MappedKind{
		FromKind: in.FromKind,
		KindName: in.KindName,
		KindLoad: in.KindLoad,
	}
}

func configFromProto(dst *config.Config, src *lpb.Config) {
	dst.RepoRoot = src.RepoRoot
}

func fileToProto(file *rule.File) *lpb.File {
	return &lpb.File{
		Pkg:       file.Pkg,
		Path:      file.Path,
		DefName:   file.DefName,
		Directive: directivesToProto(file.Directives),
		Load:      loadsToProto(file.Loads),
		Rule:      rulesToProto(file.Rules),
		Content:   file.Content,
	}
}

func rulesToProto(rules []*rule.Rule) []*lpb.Rule {
	result := make([]*lpb.Rule, len(rules))
	for i, v := range rules {
		result[i] = ruleToProto(v)
	}
	return result
}

func ruleToProto(r *rule.Rule) *lpb.Rule {
	return &lpb.Rule{
		Kind: r.Kind(),
		Name: r.Name(),
		// Fields: rule.Fields, TODO
	}
}

func loadsToProto(loads []*rule.Load) []*lpb.Load {
	result := make([]*lpb.Load, len(loads))
	for i, v := range loads {
		result[i] = loadToProto(v)
	}
	return result
}

func loadToProto(load *rule.Load) *lpb.Load {
	return &lpb.Load{
		Name: load.Name(),
		// ForceCompact: load.ForceCompact,
		// Module: load.Module,
		// Module: load.Module,
	}
}

func rulesFromProto(rules []*lpb.Rule) []*rule.Rule {
	result := make([]*rule.Rule, len(rules))
	for i, v := range rules {
		result[i] = ruleFromProto(v)
	}
	return result
}

func ruleFromProto(r *lpb.Rule) *rule.Rule {
	return rule.NewRule(r.Kind, r.Name)
}

func directivesToProto(directives []rule.Directive) []*lpb.FileDirective {
	result := make([]*lpb.FileDirective, len(directives))
	for i, v := range directives {
		result[i] = directiveToProto(v)
	}
	return result
}

func directiveToProto(directive rule.Directive) *lpb.FileDirective {
	return &lpb.FileDirective{
		Key:   directive.Key,
		Value: directive.Value,
	}
}

func configFlagSetToProto(config *config.Config, fs *flag.FlagSet) *lpb.ConfigFlagSet {
	return &lpb.ConfigFlagSet{
		Config: configToProto(config),
		Flag:   flagSetToProto(fs),
	}
}

func flagSetToProto(fs *flag.FlagSet) []*lpb.Flag {
	flags := make([]*lpb.Flag, 0)
	fs.VisitAll(func(f *flag.Flag) {
		flags = append(flags, flagToProto(f))
	})
	return flags
}

func flagToProto(f *flag.Flag) *lpb.Flag {
	return &lpb.Flag{
		Name:     f.Name,
		Usage:    f.Usage,
		Value:    f.Value.String(),
		DefValue: f.DefValue,
	}
}

func flagSetFromProto(fs *flag.FlagSet, flags []*lpb.Flag) {
	for _, f := range flags {
		flagFromProto(fs, fs.Lookup(f.Name), f)
	}
}

func flagFromProto(fs *flag.FlagSet, current *flag.Flag, in *lpb.Flag) {
	// TODO
	// if current == nil {
	// 	newFlagFromProto()
	// }
	// return &lpb.Flag{
	// 	Name:     f.Name,
	// 	Usage:    f.Usage,
	// 	Value:    f.Value,
	// 	DefValue: f.DefValue,
	// }
}
