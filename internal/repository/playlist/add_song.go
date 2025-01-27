package playlist

import (
	"context"
	"database/sql"
	"errors"

	"github.com/BelyaevEI/playlist/internal/model"
)

func (r *repo) AddSong(ctx context.Context, song *model.SongRequest) error {

	var (
		lastSong       model.Song
		firstSong      bool
		nextID, prevID sql.NullInt64
		newID          int64
	)

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
		&prevID,
		&nextID,
		&lastSong.Title,
		&lastSong.Article,
		&song.Duration); err != nil {

		// if adding song the first in playlist
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
		firstSong = true
	} else {
		if prevID.Valid {
			lastSong.Prev = prevID.Int64
		} else {
			lastSong.Prev = lastSong.ID
		}

		if nextID.Valid {
			lastSong.Next = nextID.Int64
		} else {
			lastSong.Next = 0
		}
	}

	if firstSong {
		query = `
		INSERT INTO playlist
		(user_login, title, article, duration)
		VALUES($1, $2, $3, $4)
		RETURNING id
				`

		// Insert a new song
		err = tx.QueryRowContext(ctx, query,
			song.Login,
			song.Title,
			song.Article,
			song.Duration).Scan(&newID)
		if err != nil {
			return err
		}
	} else {
		query = `
			INSERT INTO playlist
			(user_login, prev_id, next_id, title, article, duration)
			VALUES($1, $2, $3, $4, $5, $6)
			RETURNING id
				`
		// Insert a new song
		err = tx.QueryRowContext(ctx, query,
			song.Login,
			lastSong.ID,
			0, // need for connection previosly song
			song.Title,
			song.Article,
			song.Duration).Scan(&newID)
		if err != nil {
			return err
		}

		// if first song in playlist then updating is nothing
		query = `
			UPDATE playlist SET next_id = $1
			WHERE user_login = $2 AND next_id IS NULL
				`
		// Update id in previosly song
		_, err = tx.ExecContext(ctx, query, newID, song.Login)
		if err != nil {
			return err
		}

		// updating last song
		query = `
		UPDATE playlist SET next_id = NULL
		WHERE user_login = $1 AND id = $2
			`
		// Update id in previosly song
		_, err = tx.ExecContext(ctx, query, song.Login, newID)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
