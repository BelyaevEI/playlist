package playlist

import (
	"context"

	playlistRepo "github.com/BelyaevEI/playlist/internal/repository/playlist"
)

// PlaylistService represents a service for playlist entities.
type PlaylistService interface {
	StartPlayback(ctx context.Context, login string)
}

type serv struct {
	playlistRepo playlistRepo.PlaylistRepository
	userLoginCH  chan string
	usersPlaying map[string]struct{}
}

// NewService creates a new playlist service.
func NewService(playlistRepo playlistRepo.PlaylistRepository, userLoginCH chan string) PlaylistService {
	return &serv{
		playlistRepo: playlistRepo,
		userLoginCH:  userLoginCH,
		usersPlaying: make(map[string]struct{}, 0),
	}
}
