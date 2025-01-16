package auth

import (
	"context"
	"fmt"

	"github.com/BelyaevEI/playlist/internal/logger"
	"github.com/BelyaevEI/playlist/internal/model"
)

func (r *repo) SelectUser(ctx context.Context, userLogin *model.UserLogin) (model.UserLogin, error) {

	var login, passHash, secretWord string

	query := `SELECT user_login, pass_hash, secret_word FROM users WHERE user_login = $1`

	// Select user login/password and check password
	err := r.db.QueryRowContext(ctx, query, userLogin.Login).Scan(&login, &passHash, &secretWord)
	if err != nil {
		logger.Error(fmt.Sprintf("select user is failed: %s", err.Error()))
		return model.UserLogin{}, err
	}

	return model.UserLogin{Login: login,
		Password:   passHash,
		SecretWord: secretWord,
	}, nil
}
