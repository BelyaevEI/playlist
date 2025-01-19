package playlist

import (
	"context"

	"github.com/BelyaevEI/playlist/internal/model"
)

func (s *serv) NextSong(ctx context.Context, login string) {
	var cmd model.Command

	cmd.Action = "next"
	cmd.User = login

	s.actionCH <- cmd
}
