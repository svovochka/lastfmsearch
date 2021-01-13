package graph

import (
	"lastfmsearch/pkg/graph/model"
	"lastfmsearch/pkg/lastfmsearch"
)

// MapperToDomain Maps domain model to graph model
type MapperToGraph struct{}

// mapTracksList Maps domain model to graph model
func (m *MapperToGraph) mapTracksList(t []*lastfmsearch.Track) []*model.Track {
	var tracks []*model.Track
	for _, track := range t {
		tracks = append(tracks, m.mapTrack(track))
	}

	return tracks
}

// mapTrack Maps domain model to graph model
func (m *MapperToGraph) mapTrack(t *lastfmsearch.Track) *model.Track {
	return &model.Track{
		Name:      t.Name,
		URL:       t.Url,
		Listeners: t.Listeners,
		Artist:    m.mapArtist(t.Artist),
	}
}

// mapArtist Maps domain model to graph model
func (*MapperToGraph) mapArtist(a *lastfmsearch.Artist) *model.Artist {
	if a == nil {
		return nil
	}
	return &model.Artist{
		ID:      a.Id,
		Name:    a.Name,
		URL:     a.Url,
		Summary: a.Summary,
	}
}
