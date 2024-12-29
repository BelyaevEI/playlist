package playlist

import (
	"context"
	"fmt"

	"github.com/BelyaevEI/playlist/internal/logger"
	"github.com/BelyaevEI/playlist/internal/model"
)

func (s *serv) DeleteSong(ctx context.Context, song *model.SongRequest) error {
	err := s.playlistRepo.DeleteSong(ctx, song)
	if err != nil {
		logger.Error(fmt.Sprintf("delete song is failed %s", err.Error()))
		return err
	}

	return nil
}
