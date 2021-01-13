package lastfmsearch

import (
	"context"
	"fmt"
	"lastfmsearch/pkg/lastfm"
	"lastfmsearch/pkg/lastfm/method"
)

// Track Domain model
type Track struct {
	Name       string
	Url        string
	Listeners  int
	ArtistName string
	Artist     *Artist
}

// Artist Domain model
type Artist struct {
	Id      string
	Name    string
	Url     string
	Summary string
}

// Loader Search worker
type Loader struct {
	lastfmClient *lastfm.Client
	tracks       []*Track
	artists      []*Artist
}

// NewLoader Creates new loader
func NewLoader(c *lastfm.Client) *Loader {
	return &Loader{
		lastfmClient: c,
	}
}

// FindTracksByName Loads tracks list
func (l *Loader) FindTracksByName(ctx context.Context, name string, withArtists bool) ([]*Track, error) {
	result, err := method.TrackSearch(ctx, l.lastfmClient, name, 1, 100)
	if err != nil {
		return nil, fmt.Errorf("failed to get tracks list: %w", err)
	}
	//totalRows, _ := strconv.ParseInt(result.Results.OpensearchTotalResults, 0, 0)

	mapper := &MapperToDomain{}
	for _, apiTrack := range result.Results.Trackmatches.Track {
		track := mapper.mapTrack(apiTrack)
		l.tracks = append(l.tracks, track)
		if withArtists {
			result, err := method.ArtistInfo(ctx, l.lastfmClient, track.ArtistName)
			if err != nil {
				return nil, fmt.Errorf("failed to get tracks list: %w", err)
			}
			artist := mapper.mapArtist(result.Artist)
			l.artists = append(l.artists, artist)
			track.Artist = artist
		}
	}

	return l.tracks, nil
}
