package playlist

import (
	"context"
	"fmt"

	"github.com/BelyaevEI/playlist/internal/logger"
	"github.com/BelyaevEI/playlist/internal/model"
)

func (s *serv) AddSong(ctx context.Context, song *model.SongRequest) error {
	err := s.playlistRepo.AddSong(ctx, song)
	if err != nil {
		logger.Error(fmt.Sprintf("add song is failed %s", err.Error()))
		return err
	}
	return nil
}
