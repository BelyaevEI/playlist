package playlist

import "context"

func (s *serv) StartPlayback(ctx context.Context, login string) {
	_, ok := s.usersPlaying[login]
	if !ok {
		s.usersPlaying[login] = struct{}{}
		go s.playback(ctx, login)
	}
}

// playback - run playlist for user and listen command
func (s *serv) playback(ctx context.Context, login string) {

}
