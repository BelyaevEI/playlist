package auth

import (
	"context"

	"github.com/BelyaevEI/playlist/internal/converter"
	desc "github.com/BelyaevEI/playlist/pkg/auth_v1"
)

func (i *Implementation) Login(ctx context.Context, req *desc.LoginRequest) (*desc.Response, error) {

	token, err := i.authService.Login(ctx, converter.ToLoginFromDesc(req))
	if err != nil {
		return &desc.Response{RefreshToken: ""}, err
	}

	return &desc.Response{RefreshToken: token}, nil
}
