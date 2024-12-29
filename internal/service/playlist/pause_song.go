package playlist

import (
	"context"

	"github.com/BelyaevEI/playlist/internal/model"
)

func (s *serv) PauseSong(ctx context.Context, login string) {
	var cmd model.Command

	cmd.Action = "pause"
	cmd.User = login

	s.actionCH <- cmd

}
