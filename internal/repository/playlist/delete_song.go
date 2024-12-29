package playlist

import (
	"context"

	"github.com/BelyaevEI/playlist/internal/model"
)

func (r *repo) DeleteSong(ctx context.Context, song *model.SongRequest) error {

	query := `
		DELETE FROM playlist 
		WHERE user_login = $1 AND 
		title = $2 AND a
		article = $3 AND 
		duration = $4 AND
		playing = FALSE
	`
	_, err := r.db.ExecContext(ctx, query, song.Login, song.Title, song.Article, song.Duration)
	if err != nil {
		return err
	}
	return nil
}
