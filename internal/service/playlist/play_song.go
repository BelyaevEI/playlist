package playlist

import (
	"context"

	"github.com/BelyaevEI/playlist/internal/model"
)

func (s *serv) PlaySong(ctx context.Context, login string) {
	var cmd model.Command

	cmd.Action = "play"
	cmd.User = login

	s.userLoginCH <- login
	s.actionCH <- cmd

}
