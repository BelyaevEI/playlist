package auth

import (
	authService "github.com/BelyaevEI/playlist/internal/service/auth"
	desc "github.com/BelyaevEI/playlist/pkg/auth_v1"
)

type Implementation struct {
	desc.UnimplementedAuthV1Server
	authService authService.AuthService
}

// NewImplementation creates a new auth API implementation.
func NewImplementation(authService authService.AuthService) *Implementation {

	return &Implementation{
		authService: authService,
	}
}
