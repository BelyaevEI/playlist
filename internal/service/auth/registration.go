package auth

import (
	"context"
	"errors"

	"github.com/BelyaevEI/playlist/internal/logger"
	"github.com/BelyaevEI/playlist/internal/model"
)

func (s *serv) Registration(ctx context.Context, userRegistration *model.UserRegistration) (string, error) {

	// Check password = confirm password
	if userRegistration.Password != userRegistration.ConfirmPassword {
		logger.Info("password and confirm not equal")
		return "", errors.New("password and confirm not equal")
	}

	// Create user
	err := s.authRepo.CreateUser(ctx, *userRegistration)
	if err != nil {
		logger.Debug(err.Error())
		return "", err
	}

	return "", nil
}
