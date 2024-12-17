package playlist

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/BelyaevEI/playlist/internal/logger"
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

		// given action command from user
		case cmd := <-s.actionCH:
			switch cmd.Action {
			case "play":
				if cmd.User == playlist.User {
					playlist.Playing = true

					// if current song is nil then start playlist
					// or playlist is empty
					if playlist.Current == nil {
						song, err := s.playlistRepo.GetFirstSongOfUser(ctx, playlist.User)
						if err != nil {
							logger.Error(fmt.Sprintf("get first song of user is failed: %v", err))
							break
						}

						//play the first song of playlist
						playlist.Current = &song
					}

				}
			}
		default:

			if playlist.Playing && playlist.Current != nil {

				// current song finished play and play next song
				if playlist.Current.TimeSong == 0 {
					song, err := s.playlistRepo.GetNextSongOfUser(
						ctx,
						playlist.User,
						playlist.Current.ID,
						playlist.Current.Next,
					)
					if err != nil {

						// if songs is finished in playlist then get first song and playback
						if err == sql.ErrNoRows {
							song, err := s.playlistRepo.GetFirstSongOfUser(ctx, playlist.User)
							if err != nil {
								logger.Error(fmt.Sprintf("get first song of user is failed: %v", err))
								break
							}

							playlist.Current = &song
						}
						logger.Error(fmt.Sprintf("get next song of user is failed: %v", err))
					}

					playlist.Current = &song
				}

				// play user  current song
				playlist.Current.TimeSong -= 1

				logger.Info(
					fmt.Sprintf("playing now: %s - %s, time %d : %d : %d ",
						playlist.Current.Title,
						playlist.Current.Article,
						int(playlist.Current.TimeSong.Hours()),
						int(playlist.Current.TimeSong.Minutes())%60,
						int(playlist.Current.TimeSong.Seconds())%60),
				)

			}
		}
	}
}
