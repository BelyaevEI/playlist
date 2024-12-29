package playlist

import (
	"context"

	"github.com/BelyaevEI/playlist/internal/model"
)

func (r *repo) AddSong(ctx context.Context, song *model.SongRequest) error {

	var lastSong model.Song

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := `
		SELECT ID, prev_id, next_id, title, article, duration
		FROM playlist
		WHERE next_id IS NULL
		AND user_login = $1
	`

	// Get last song for insert new song in tail
	row := tx.QueryRowContext(ctx, query, song.Login)
	if err := row.Scan(
		&lastSong.ID,
		&lastSong.Prev,
		&lastSong.Next,
		&lastSong.Title,
		&lastSong.Article,
		&song.Duration); err != nil {

		return err
	}

	query = `
	INSERT INTO playlist
	(user_login, prev_id, title, article, duartion)
	VALUES($1, $2, $3, $4, $5)
`

	// Insert a new song
	var newID int64
	err = tx.QueryRowContext(ctx, query,
		song.Login,
		lastSong.ID,
		song.Title,
		song.Article,
		song.Duration).Scan(&newID)
	if err != nil {
		return err
	}

	query = `
		UPDATE playlist SET next_id = $1
		WHERE user_login = $2 AND next_id IS NULL
	`
	// Update id in previosly song
	_, err = tx.ExecContext(ctx, query, newID, song.Login)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
