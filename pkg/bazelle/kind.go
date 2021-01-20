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

func (x *rpcLang) Kinds() map[string]rule.KindInfo {
	infos := make()
	return map[string]rule.KindInfo{
		"x_library": {},
	}
}

func kindInfosFromProto(Kinds []*pb.KindInfo) []rule.KindInfo {
	infos := make([]rule.KindInfo, len(Kinds))
	for i, Kind := range Kinds {
		infos[i] = kindInfoFromProto(Kind)
	}
	return infos
}

func kindInfoFromProto(Kind *pb.KindInfo) rule.KindInfo {
	return rule.KindInfo{Kind.Name, Kind.Symbols, Kind.After}
}
