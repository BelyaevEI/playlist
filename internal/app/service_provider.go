package app

import (
	"context"

	"github.com/BelyaevEI/test-assignment/internal/api/auth"
	"github.com/BelyaevEI/test-assignment/internal/config"
	authRepo "github.com/BelyaevEI/test-assignment/internal/repository/auth"
	authService "github.com/BelyaevEI/test-assignment/internal/service/auth"
)

type serviceProvider struct {
	config         config.Config
	authImpl       *auth.Implementation
	authService    authService.AuthService
	authRepository authRepo.AuthRepository
}

func newServiceProvider(cfg config.Config) *serviceProvider {
	return &serviceProvider{config: cfg}
}

// AuthImlp - implementation auth api layer
func (s *serviceProvider) AuthImpl(ctx context.Context) *auth.Implementation {
	if s.authImpl == nil {
		s.authImpl = auth.NewImplementation(s.AuthService(ctx))
	}

	return s.authImpl
}

// AuthService - implementation service layer
func (s *serviceProvider) AuthService(ctx context.Context) authService.AuthService {
	if s.authService == nil {
		s.authService = authService.NewService(s.AuthRepository(ctx))
	}

	return s.authService
}

// AuthRepositiry - implementation repository layer
func (s *serviceProvider) AuthRepository(ctx context.Context) authRepo.AuthRepository {
	if s.authRepository == nil {
		s.authRepository = authRepo.NewRepository(ctx, s.config.DSN())
	}

	return s.authRepository
}
