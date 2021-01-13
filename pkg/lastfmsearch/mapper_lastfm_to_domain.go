package lastfmsearch

import (
	"lastfmsearch/pkg/lastfm/method"
	"strconv"
)

// MapperToDomain Maps api model to domain model
type MapperToDomain struct{}

// mapTrack Maps api model to domain model
func (m *MapperToDomain) mapTrack(t *method.TrackItem) *Track {
	listeners, _ := strconv.ParseInt(t.Listeners, 0, 0)
	return &Track{
		Name:       t.Name,
		Url:        t.Url,
		Listeners:  int(listeners),
		ArtistName: t.Artist,
		Artist:     nil,
	}
}

// mapArtist Maps api model to domain model
func (m *MapperToDomain) mapArtist(a *method.Artist) *Artist {
	return &Artist{
		Id:      a.Id,
		Name:    a.Name,
		Url:     a.Url,
		Summary: a.Bio.Summary,
	}
}
