package method

import (
	"context"
	"fmt"
	"lastfmsearch/pkg/lastfm"
	"regexp"
)

// ArtistInfoResp API response
type ArtistInfoResp struct {
	Artist *Artist `json:"artist"`
}

// Artist Domain model
type Artist struct {
	Id   string `json:"mbid"`
	Name string `json:"name"`
	Url  string `json:"url"`
	Bio  Bio    `json:"bio"`
}

// Bio Domain model
type Bio struct {
	Summary string `json:"summary"`
}

// ArtistInfo Returns artist info
func ArtistInfo(ctx context.Context, client *lastfm.Client, name string) (*ArtistInfoResp, error) {
	regex := regexp.MustCompile("\\s+")
	name = regex.ReplaceAllString(name, "+")
	query := fmt.Sprintf(
		"%s/?api_key=%s&method=artist.getinfo&artist=%s&format=json",
		client.Endpoint,
		client.ApiKey,
		name,
	)
	var result *ArtistInfoResp
	err := client.Query(ctx, query, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to get artist info: %w", err)
	}

	return result, nil
}
