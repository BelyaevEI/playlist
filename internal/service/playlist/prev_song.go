package playlist

import (
	"context"

	"github.com/BelyaevEI/playlist/internal/model"
)

func (s *serv) PrevSong(ctx context.Context, login string) {
	var cmd model.Command

	cmd.Action = "prev"
	cmd.User = login

	s.actionCH <- cmd

}
