package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"

	"github.com/BelyaevEI/playlist/internal/logger"
	"github.com/BelyaevEI/playlist/internal/model"

	"golang.org/x/crypto/bcrypt"
)

func (s *serv) Registration(ctx context.Context, userRegistration *model.UserRegistration) (string, error) {

	// Check password = confirm password
	if userRegistration.Password != userRegistration.ConfirmPassword {
		logger.Info("password and confirm password not equal")
		return "", errors.New("password and confirm password not equal")
	}

	// Check login is unique
	err := s.authRepo.CheckLoginUnique(ctx, userRegistration.Login)
	if err != nil {
		logger.Info("login is not unique")
		return "", errors.New("login is not unique")
	}

	// Generate secret word for jwt
	secret := make([]byte, 12)

	_, err = rand.Read(secret)
	if err != nil {
		logger.Debug(err.Error())
		return "", err
	}

	secretWord := base64.URLEncoding.EncodeToString(secret)

	// Hashing user password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRegistration.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// Create user
	err = s.authRepo.CreateUser(ctx, userRegistration.Login, string(hashedPassword), secretWord)
	if err != nil {
		logger.Debug(err.Error())
		return "", err
	}

	// Create a new jwt for user
	// if generating token is failed then user must be log in with login and password
	token, err := generateToken(userRegistration.Login, secretWord)
	if err != nil {
		logger.Debug(err.Error())
	}

	return token, nil
}
