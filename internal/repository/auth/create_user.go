package auth

import (
	"context"
	"fmt"

	"github.com/BelyaevEI/playlist/internal/logger"
)

func (r *repo) CreateUser(ctx context.Context, login, hashedPassword, secretKey string) error {
	query := `
		INSERT INTO users (user_login, pass_hash, secret_word)
		VALUES ($1, $2, $3))
	`

	_, err := r.db.ExecContext(ctx, query, login, hashedPassword, secretKey)
	if err != nil {
		logger.Error(fmt.Sprintf("create user is failed: %s", err))
		return err
	}

	return nil
}

func (r *repo) CheckLoginUnique(ctx context.Context, login string) error {
	var notUnique string
	query := `SELECT user_login FROM users WHERE user_login = $1`

	err := r.db.QueryRowContext(ctx, query, login).Scan(&notUnique)
	if err != nil {
		logger.Error(fmt.Sprintf("check unique login: %s", login))
		return err
	}

	return nil
}
