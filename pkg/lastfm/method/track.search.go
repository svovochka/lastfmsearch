package method

import (
	"context"
	"fmt"
	"lastfmsearch/pkg/lastfm"
	"net/url"
)

// TrackSearchResp API response
type TrackSearchResp struct {
	Results *Results `json:"results"`
}

// Results Domain model
type Results struct {
	OpensearchTotalResults string        `json:"opensearch:totalResults"`
	Trackmatches           *Trackmatches `json:"trackmatches"`
}

// Trackmatches Domain model
type Trackmatches struct {
	Track []*TrackItem `json:"track"`
}

// Track Domain model
type TrackItem struct {
	Name      string `json:"name"`
	Artist    string `json:"artist"`
	Url       string `json:"url"`
	Listeners string `json:"listeners"`
}

// ArtistInfo Returns artist info
func TrackSearch(ctx context.Context, client *lastfm.Client, name string, page int, limit int) (*TrackSearchResp, error) {
	name = url.QueryEscape(name)
	query := fmt.Sprintf(
		"%s/?api_key=%s&method=track.search&track=%s&page=%d&limit=%d&format=json",
		client.Endpoint,
		client.ApiKey,
		name,
		page,
		limit,
	)
	var result *TrackSearchResp
	err := client.Query(ctx, query, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to get tracks list: %w", err)
	}

	return result, nil
}
