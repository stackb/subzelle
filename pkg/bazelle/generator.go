package subzelle

import (
	"context"

	"github.com/bazelbuild/bazel-gazelle/language"
	"github.com/bazelbuild/bazel-gazelle/rule"

	lpb "github.com/stackb/subzelle/language"
)

func (b *subzelle) Kinds() map[string]rule.KindInfo {
	response, err := b.languageClient.Kinds(context.Background(), &lpb.KindsRequest{})
	if err != nil {
		fatalError(err)
	}

	result := make(map[string]rule.KindInfo)
	for k, v := range response.Kinds {
		result[k] = kindInfoFromProto(v)
	}

	return result
}

func (b *subzelle) Loads() []rule.LoadInfo {
	response, err := b.languageClient.Loads(context.Background(), &lpb.LoadsRequest{})
	if err != nil {
		fatalError(err)
	}

	return loadInfosFromProto(response.Load)
}

func (b *subzelle) GenerateRules(args language.GenerateArgs) language.GenerateResult {
	request := generateArgsToProto(args)

	response, err := b.languageClient.GenerateRules(context.Background(), request)
	if err != nil {
		fatalError(err)
	}

	return generateResultFromProto(response)
}
