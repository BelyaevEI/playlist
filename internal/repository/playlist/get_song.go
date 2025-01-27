package playlist

import (
	"context"
	"database/sql"

	"github.com/BelyaevEI/playlist/internal/model"
)

func (r *repo) GetFirstSongOfUser(ctx context.Context, login string) (model.Song, error) {

	var (
		song                   model.Song
		currID, nextID, prevID sql.NullInt64
	)

	tx, err := r.db.Begin()
	if err != nil {
		return model.Song{}, err
	}

	defer tx.Rollback()

	// get song
	query := `
		SELECT id, playing, prev_id, next_id, title, article, duration
		FROM playlist WHERE user_login = $1 AND prev_id IS NULL
	`

	row := tx.QueryRowContext(ctx, query, login)
	if err := row.Scan(
		&currID,
		&song.Playnig,
		&prevID,
		&nextID,
		&song.Title,
		&song.Article,
		&song.Duration); err != nil {

		return model.Song{}, err
	}

	if prevID.Valid {
		song.Prev = prevID.Int64
	}
	if nextID.Valid {
		song.Next = nextID.Int64
	}
	if currID.Valid {
		song.ID = currID.Int64
	}
	// update song for defend changes
	query = `
		UPDATE playlist SET playing = $1
		WHERE user_login = $2 AND prev_id IS NULL
	`

	_, err = tx.ExecContext(ctx, query, true, login)
	if err != nil {
		return model.Song{}, err
	}

	err = tx.Commit()
	if err != nil {
		return model.Song{}, err
	}

	return song, nil
}

func (r *repo) GetNextSongOfUser(ctx context.Context, login string, currID, nextID int64) (model.Song, error) {

	var (
		song          model.Song
		cID, nID, pID sql.NullInt64
	)

	tx, err := r.db.Begin()
	if err != nil {
		return model.Song{}, err
	}

	// checking next ID if equal zero then need given first song
	if nextID == 0 {
		return model.Song{}, sql.ErrNoRows
	}

	defer tx.Rollback()

	query := `
		SELECT id, playing, prev_id, next_id, title, article, duration
		FROM playlist WHERE user_login = $1 AND id = $2
	`

	row := tx.QueryRowContext(ctx, query, login, nextID)
	if err := row.Scan(
		&cID,
		&song.Playnig,
		&pID,
		&nID,
		&song.Title,
		&song.Article,
		&song.Duration); err != nil {

		return model.Song{}, err
	}

	if cID.Valid {
		song.ID = cID.Int64
	}
	if pID.Valid {
		song.Prev = pID.Int64
	}
	if nID.Valid {
		song.Next = nID.Int64
	}

	// update song for defend changes
	query = `
		UPDATE playlist SET playing = $1
		WHERE user_login = $2 AND id = $3
	`

	_, err = tx.ExecContext(ctx, query, false, login, currID)
	if err != nil {
		return model.Song{}, err
	}

	// update song for defend changes
	query = `
		UPDATE playlist SET playing = $1
		WHERE user_login = $2 AND id = $3
	`

	_, err = tx.ExecContext(ctx, query, true, login, nextID)
	if err != nil {
		return model.Song{}, err
	}

	err = tx.Commit()
	if err != nil {
		return model.Song{}, err
	}

	return song, nil
}

func (r *repo) GetPrevSongOfUser(ctx context.Context, login string, currID, prevID int64) (model.Song, error) {

	var (
		song          model.Song
		cID, nID, pID sql.NullInt64
	)

	tx, err := r.db.Begin()
	if err != nil {
		return model.Song{}, err
	}

	// checking previosly ID if equal zero then need given first song
	if prevID == 0 {
		return model.Song{}, sql.ErrNoRows
	}

	defer tx.Rollback()

	query := `
		SELECT id, playing, prev_id, next_id, title, article, duration
		FROM playlist WHERE user_login = $1 AND id = $2
	`

	row := tx.QueryRowContext(ctx, query, login, prevID)
	if err := row.Scan(
		&cID,
		&song.Playnig,
		&pID,
		&nID,
		&song.Title,
		&song.Article,
		&song.Duration); err != nil {

		return model.Song{}, err
	}

	if cID.Valid {
		song.ID = cID.Int64
	}
	if pID.Valid {
		song.Prev = pID.Int64
	}
	if nID.Valid {
		song.Next = nID.Int64
	}

	// update song for defend changes
	query = `
		UPDATE playlist SET playing = $1
		WHERE user_login = $2 AND id = $3
	`

	_, err = tx.ExecContext(ctx, query, false, login, currID)
	if err != nil {
		return model.Song{}, err
	}

	// update song for defend changes
	query = `
		UPDATE playlist SET playing = $1
		WHERE user_login = $2 AND id = $3
	`

	_, err = tx.ExecContext(ctx, query, true, login, prevID)
	if err != nil {
		return model.Song{}, err
	}

	err = tx.Commit()
	if err != nil {
		return model.Song{}, err
	}

	return song, nil
}
