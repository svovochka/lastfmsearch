package graph

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"lastfmsearch/pkg/graph/model"
)

type Resolver struct{}

func (r *Resolver) doFindTracksByTitle(ctx context.Context, title *string) ([]*model.Track, error) {
	//preloads := getPreloads(ctx)
	return nil, errors.New("something fucked up")
}

func getPreloads(ctx context.Context) []string {
	return getNestedPreloads(
		graphql.GetOperationContext(ctx),
		graphql.CollectFieldsCtx(ctx, nil),
		"",
	)
}

func getNestedPreloads(ctx *graphql.OperationContext, fields []graphql.CollectedField, prefix string) (preloads []string) {
	for _, column := range fields {
		prefixColumn := getPreloadString(prefix, column.Name)
		preloads = append(preloads, prefixColumn)
		preloads = append(preloads, getNestedPreloads(ctx, graphql.CollectFields(ctx, column.Selections, nil), prefixColumn)...)
	}
	return
}

func getPreloadString(prefix, name string) string {
	if len(prefix) > 0 {
		return prefix + "." + name
	}
	return name
}
