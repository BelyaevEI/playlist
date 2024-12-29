package playlist

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	"github.com/BelyaevEI/playlist/internal/logger"
	"github.com/BelyaevEI/playlist/internal/model"
)

func (s *serv) CloseActionCH() {
	close(s.actionCH)
}

func (s *serv) StartPlayback(ctx context.Context, login string, wg *sync.WaitGroup) {
	_, ok := s.usersPlaying[login]
	if !ok {
		s.mu.Lock()
		s.usersPlaying[login] = struct{}{}
		s.mu.Unlock()
		s.playback(ctx, login, wg)
	}
}

// playback - run playlist of user and listen command
func (s *serv) playback(ctx context.Context, login string, wg *sync.WaitGroup) {

	defer wg.Done()

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
							logger.Error(fmt.Sprintf("getting first song of user is failed: %v", err))
							break
						}

						//playing the first song of playlist
						playlist.Current = &song
					}
				}

			case "pause":
				// if getting command pause then stop playing song
				if cmd.User == playlist.User {
					playlist.Playing = false
				}

			case "next":
				// if getting command next song then switch playing song
				if cmd.User == playlist.User {
					playlist.Playing = true

					song, err := s.playlistRepo.GetNextSongOfUser(
						ctx,
						playlist.User,
						playlist.Current.ID,
						playlist.Current.Next,
					)
					if err != nil {

						// if songs is finished in playlist then getting first song and playback
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

			case "prev":
				// if getting command prev song then switch playing song
				if cmd.User == playlist.User {
					playlist.Playing = true

					song, err := s.playlistRepo.GetPrevSongOfUser(
						ctx,
						playlist.User,
						playlist.Current.ID,
						playlist.Current.Next,
					)
					if err != nil {
						// if songs is finished in playlist then getting first song and playback
						if err == sql.ErrNoRows {
							song, err := s.playlistRepo.GetFirstSongOfUser(ctx, playlist.User)
							if err != nil {
								logger.Error(fmt.Sprintf("get first song of user is failed: %v", err))
								break
							}

							playlist.Current = &song
						}
						logger.Error(fmt.Sprintf("get prev song of user is failed: %v", err))
					}

					playlist.Current = &song
				}
			}
		default:

			if playlist.Playing && playlist.Current != nil {

				// current song finished play and playing next song
				if playlist.Current.Duration == 0 {
					song, err := s.playlistRepo.GetNextSongOfUser(
						ctx,
						playlist.User,
						playlist.Current.ID,
						playlist.Current.Next,
					)
					if err != nil {

						// if songs is finished in playlist then getting first song and playback
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

				// play user current song
				playlist.Current.Duration -= 1

				logger.Info(
					fmt.Sprintf("playing now: %s - %s, time %d : %d : %d",
						playlist.Current.Title,
						playlist.Current.Article,
						int(playlist.Current.Duration.Hours()),
						int(playlist.Current.Duration.Minutes())%60,
						int(playlist.Current.Duration.Seconds())%60),
				)

			}
		}
	}
}
