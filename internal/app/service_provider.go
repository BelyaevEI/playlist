package app

import (
	"context"
	"sync"

	"github.com/BelyaevEI/playlist/internal/api/auth"
	"github.com/BelyaevEI/playlist/internal/api/playlist"
	"github.com/BelyaevEI/playlist/internal/config"
	"github.com/BelyaevEI/playlist/internal/interceptor"
	"github.com/BelyaevEI/playlist/internal/logger"
	authRepo "github.com/BelyaevEI/playlist/internal/repository/auth"
	playlistRepo "github.com/BelyaevEI/playlist/internal/repository/playlist"
	authService "github.com/BelyaevEI/playlist/internal/service/auth"
	playlistService "github.com/BelyaevEI/playlist/internal/service/playlist"
)

type serviceProvider struct {
	config          config.Config
	authInterceptor interceptor.AuthInterceptor

	authImpl       *auth.Implementation
	authService    authService.AuthService
	authRepository authRepo.AuthRepository

	plImpl       *playlist.Implementation
	plService    playlistService.PlaylistService
	plRepository playlistRepo.PlaylistRepository

	userLoginCH chan string
}

func newServiceProvider(cfg config.Config, authInterceptor interceptor.AuthInterceptor) *serviceProvider {
	return &serviceProvider{
		config:          cfg,
		userLoginCH:     make(chan string),
		authInterceptor: authInterceptor,
	}
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

// PlaylistRepositiry - implementation repository layer
func (s *serviceProvider) PlaylistRepository(ctx context.Context) playlistRepo.PlaylistRepository {
	if s.plRepository == nil {
		s.plRepository = playlistRepo.NewRepository(ctx, s.config.DSN())
	}

	return s.plRepository
}

// PlaylistService - implementation service layer
func (s *serviceProvider) PlaylistService(ctx context.Context, userLoginCH chan string) playlistService.PlaylistService {
	if s.plService == nil {
		s.plService = playlistService.NewService(s.PlaylistRepository(ctx), userLoginCH)
	}

	return s.plService
}

// PlaylistImlp - implementation auth api layer
func (s *serviceProvider) PlaylistImpl(ctx context.Context) *playlist.Implementation {
	if s.plImpl == nil {
		s.plImpl = playlist.NewImplementation(s.PlaylistService(ctx, s.userLoginCH))
	}

	return s.plImpl
}

// StartPlayback - start goroutine for user playlist
func (s *serviceProvider) StartPlayback(ctx context.Context) {
	wg := sync.WaitGroup{}

	// Given login user for playback playlist
	for login := range s.userLoginCH {
		wg.Add(1)
		go s.plService.StartPlayback(ctx, login, &wg)
	}

	wg.Wait()
	logger.Info("stop running playback")
}

func (s *serviceProvider) CloseActionCH() {
	s.plService.CloseActionCH()
}

// Initializating authentification interceptor
func (s *serviceProvider) InitAuthInterceptor(ctx context.Context) interceptor.AuthInterceptor {
	if s.authInterceptor == nil {
		s.authInterceptor = interceptor.NewAuthInterceptor(ctx, s.config.DSN())
	}

	return s.authInterceptor
}
