package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"lastfmsearch/pkg/graph/generated"
	"lastfmsearch/pkg/graph/model"
)

func (r *queryResolver) FindTracksByName(ctx context.Context, name *string) ([]*model.Track, error) {
	return r.doFindTracksByName(ctx, name)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
