package auth

import (
	"context"

	"github.com/BelyaevEI/test-assignment/internal/model"
	authRepo "github.com/BelyaevEI/test-assignment/internal/repository/auth"
)

// AuthService represents a service for auth entities.
type AuthService interface {
	Login(ctx context.Context, userLogin *model.UserLogin) (string, error)
	Registration(ctx context.Context, userRegistration *model.UserRegistration) (string, error)
}

type serv struct {
	authRepo authRepo.AuthRepository
}

// NewService creates a new auth service.
func NewService(authRepo authRepo.AuthRepository) AuthService {
	return &serv{authRepo: authRepo}
}
