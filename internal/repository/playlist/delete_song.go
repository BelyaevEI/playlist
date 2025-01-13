package playlist

import (
	"context"
	"database/sql"

	"github.com/BelyaevEI/playlist/internal/model"
)

func (r *repo) DeleteSong(ctx context.Context, song *model.SongRequest) error {

	var currID, nextID, prevID sql.NullInt64

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	//Need select deleting song and change links
	query := `
		SELECT id, prev_id, next_id
		FROM playlist
		WHERE user_login = $1
		AND title = $2
		AND article = $3
		`

	row := tx.QueryRowContext(ctx, query, song.Login, song.Title, song.Article)
	if err := row.Scan(&currID, &prevID, &nextID); err != nil {
		return err
	}

	query = `
		DELETE FROM playlist 
		WHERE user_login = $1 AND 
		title = $2 AND 
		article = $3 AND 
		duration = $4 AND
		playing = FALSE
	`
	_, err = tx.ExecContext(ctx, query, song.Login, song.Title, song.Article, song.Duration)
	if err != nil {
		return err
	}

	//Updating links for prev and next song
	if prevID.Valid {
		// Update id in previosly song
		// if current nextID is not null then update previosly song nextID
		if nextID.Valid {
			query = `
				UPDATE playlist SET next_id = $1
				WHERE user_login = $2
				AND id = $3
				`
			_, err = tx.ExecContext(ctx, query, nextID.Int64, song.Login, prevID.Int64)
			if err != nil {
				return err
			}

			query = `
				UPDATE playlist SET prev_id = $1
				WHERE user_login = $2
				AND id = $3
				`
			_, err = tx.ExecContext(ctx, query, prevID.Int64, song.Login, nextID.Int64)
			if err != nil {
				return err
			}

		} else {
			// else this last song in playlist
			query = `
				UPDATE playlist SET next_id = NULL
				WHERE user_login = $1
				AND id = $2
				`
			// Update id in previosly song
			_, err = tx.ExecContext(ctx, query, song.Login, prevID.Int64)
			if err != nil {
				return err
			}
		}
	} else {
		if nextID.Valid {
			query = `
				UPDATE playlist SET prev_id = NULL
				WHERE user_login = $1
				AND id = $2
				`
			// Update id in previosly song
			_, err = tx.ExecContext(ctx, query, song.Login, nextID.Int64)
			if err != nil {
				return err
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
