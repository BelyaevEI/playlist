package playlist

import "context"

func (s *serv) StartPlayback(ctx context.Context, login string) {
	usersPlayNow := make(map[string]struct{}, 0)

	_, ok := usersPlayNow[login]
	if !ok {
		usersPlayNow[login] = struct{}{}

	}
}
