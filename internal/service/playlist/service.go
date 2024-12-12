package playlist

import (
	playlistRepo "github.com/BelyaevEI/playlist/internal/repository/playlist"
)

// PlaylistService represents a service for playlist entities.
type PlaylistService interface {
}

type serv struct {
	playlistRepo playlistRepo.PlaylistRepository
}

// NewService creates a new playlist service.
func NewService(playlistRepo playlistRepo.PlaylistRepository) PlaylistService {
	return &serv{playlistRepo: playlistRepo}
}
