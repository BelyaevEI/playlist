package playlist

import (
	playlistService "github.com/BelyaevEI/playlist/internal/service/playlist"
	descPlaylist "github.com/BelyaevEI/playlist/pkg/playlist_v1"
)

type Implementation struct {
	descPlaylist.UnimplementedPlaylistV1Server
	playlistService playlistService.PlaylistService
}

// NewImplementation creates a new playlist API implementation.
func NewImplementation(playlistService playlistService.PlaylistService) *Implementation {

	return &Implementation{
		playlistService: playlistService,
	}
}
