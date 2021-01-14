package lastfmsearch

import (
	"context"
	"fmt"
	"lastfmsearch/pkg/lastfm"
	"lastfmsearch/pkg/lastfm/method"
	"math"
	"strconv"
	"sync"
)

const perPage = 100

// TODO This limitation was added in dev purposes because the number of results could be astronomic :)
const maxPages = 2

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
	tracksMu     sync.Mutex
	tracks       []*Track
	artistsMu    sync.Mutex
	artists      []*Artist
}

// NewLoader Creates new loader
func NewLoader(c *lastfm.Client) *Loader {
	return &Loader{
		lastfmClient: c,
	}
}

// FindTracksByName Loads tracks list
func (l *Loader) FindTracksByName(ctx context.Context, name string, withArtist bool) ([]*Track, error) {
	errCh := make(chan error)
	doneCh := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(1)
	go loadTracksPage(ctx, &wg, errCh, l, 1, name, withArtist)
	go func() {
		wg.Wait()
		close(doneCh)
	}()

	select {
	case <-doneCh:
		break
	case err := <-errCh:
		return nil, err
	}

	return l.tracks, nil
}

// loadTracksPage Loads tracks page
func loadTracksPage(
	ctx context.Context,
	wg *sync.WaitGroup,
	errCh chan error,
	l *Loader,
	page int,
	name string,
	withArtist bool,
) {
	defer wg.Done()

	result, err := method.TrackSearch(ctx, l.lastfmClient, name, page, perPage)
	if err != nil {
		errCh <- fmt.Errorf("failed to get tracks list: %w", err)
		return
	}

	mapper := &MapperToDomain{}
	for _, apiTrack := range result.Results.Trackmatches.Track {
		track := mapper.mapTrack(apiTrack)
		l.tracksMu.Lock()
		l.tracks = append(l.tracks, track)
		l.tracksMu.Unlock()
		if withArtist {
			l.artistsMu.Lock()
			for _, artist := range l.artists {
				if artist.Name == track.ArtistName {
					track.Artist = artist
					break
				}
			}
			l.artistsMu.Unlock()
			if track.Artist == nil {
				wg.Add(1)
				go loadArtistForTrack(ctx, wg, errCh, l, track)
			}
		}
	}

	nextPage := page + 1
	totalRows, _ := strconv.ParseInt(result.Results.OpensearchTotalResults, 0, 0)
	totalPages := int(math.Ceil(float64(totalRows) / float64(perPage)))
	if nextPage > totalPages || nextPage > maxPages {
		return
	}

	wg.Add(1)
	go loadTracksPage(ctx, wg, errCh, l, nextPage, name, withArtist)
}

func loadArtistForTrack(ctx context.Context, wg *sync.WaitGroup, errCh chan error, l *Loader, track *Track) {
	defer wg.Done()

	result, err := method.ArtistInfo(ctx, l.lastfmClient, track.ArtistName)
	if err != nil {
		errCh <- fmt.Errorf("failed to get tracks list: %w", err)
		return
	}

	mapper := &MapperToDomain{}
	artist := mapper.mapArtist(result.Artist)
	l.artistsMu.Lock()
	l.artists = append(l.artists, artist)
	l.artistsMu.Unlock()
	track.Artist = artist
}
