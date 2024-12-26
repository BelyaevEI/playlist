package playlist

import (
	"context"
	"sync"

	"github.com/BelyaevEI/playlist/internal/model"
	playlistRepo "github.com/BelyaevEI/playlist/internal/repository/playlist"
)

// PlaylistService represents a service for playlist entities.
type PlaylistService interface {
	StartPlayback(ctx context.Context, login string, wg *sync.WaitGroup)
	AddSong(ctx context.Context, song *model.SongRequest) error
	CloseActionCH()
}

type serv struct {
	playlistRepo playlistRepo.PlaylistRepository
	userLoginCH  chan string
	actionCH     chan model.Command
	usersPlaying map[string]struct{}
	mu           sync.Mutex
}

// NewService creates a new playlist service.
func NewService(playlistRepo playlistRepo.PlaylistRepository, userLoginCH chan string) PlaylistService {
	return &serv{
		playlistRepo: playlistRepo,
		userLoginCH:  userLoginCH,
		usersPlaying: make(map[string]struct{}, 0),
		actionCH:     make(chan model.Command),
	}
}
