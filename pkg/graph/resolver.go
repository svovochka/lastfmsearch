package graph

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"lastfmsearch/pkg/graph/model"
	"lastfmsearch/pkg/lastfm"
	"lastfmsearch/pkg/lastfmsearch"
)

// NewClient Create new client
func NewResolver(lastfmClient *lastfm.Client) *Resolver {
	return &Resolver{
		lastfmClient: lastfmClient,
	}
}

type Resolver struct {
	lastfmClient *lastfm.Client
}

func (r *Resolver) doFindTracksByName(ctx context.Context, name *string) ([]*model.Track, error) {
	withArtist := false
	for _, field := range graphql.CollectFieldsCtx(ctx, nil) {
		if field.Name == "artist" {
			withArtist = true
		}
	}
	loader := lastfmsearch.NewLoader(r.lastfmClient)
	tracks, err := loader.FindTracksByName(ctx, *name, withArtist)
	if err != nil {
		return nil, fmt.Errorf("failed to execute http request: %w", err)
	}
	mapper := &MapperToGraph{}

	return mapper.mapTracksList(tracks), nil
}
