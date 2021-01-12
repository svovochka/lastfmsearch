package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"lastfmsearch/pkg/graph/generated"
	"lastfmsearch/pkg/graph/model"
)

func (r *queryResolver) FindTracksByTitle(ctx context.Context, title *string) ([]*model.Track, error) {
	return r.doFindTracksByTitle(ctx, title)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
