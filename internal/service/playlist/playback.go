package playlist

import (
	"context"

	"github.com/BelyaevEI/playlist/internal/model"
)

func (s *serv) StartPlayback(ctx context.Context, login string) {
	_, ok := s.usersPlaying[login]
	if !ok {
		s.mu.Lock()
		s.usersPlaying[login] = struct{}{}
		s.mu.Unlock()
		s.playback(ctx, login)
	}
}

// playback - run playlist of user and listen command
func (s *serv) playback(ctx context.Context, login string) {

	// playlist of user
	playlist := &model.Playlist{
		User: login,
	}

	for {
		select {
		case cmd := <-s.actionCH:
			switch cmd.Action {
			case "play":
				if cmd.User == playlist.User {
					playlist.Playing = true
					// if current song is nil then start playlist
					// or playlist is empty
					if playlist.Current == nil {

					}
					//play the first song of playlist
				}
			}
		default:
			if playlist.Playing && playlist.Current != nil {
			}
		}
	}
}
