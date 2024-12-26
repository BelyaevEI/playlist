package auth

import (
	"context"
	"fmt"

	jwtutils "github.com/BelyaevEI/playlist/internal/jwt_utils"
	"github.com/BelyaevEI/playlist/internal/logger"
	"github.com/BelyaevEI/playlist/internal/model"

	"golang.org/x/crypto/bcrypt"
)

func (s *serv) Login(ctx context.Context, userLogin *model.UserLogin) (string, error) {

	selectUser, err := s.authRepo.SelectUser(ctx, userLogin)
	if err != nil {
		return "", err
	}

	// Compare entered password and password from database
	err = bcrypt.CompareHashAndPassword([]byte(selectUser.Password), []byte(userLogin.Password))
	if err != nil {
		logger.Debug(fmt.Sprintf("not compare hash and password: %s", err.Error()))
		return "", err
	}

	// Create a new jwt for user
	token, err := jwtutils.GenerateToken(selectUser.Login, selectUser.SecretWord)
	if err != nil {
		logger.Debug(err.Error())
	}

	return token, nil
}
