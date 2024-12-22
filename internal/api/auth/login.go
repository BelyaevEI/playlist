package auth

import (
	"context"

	"github.com/BelyaevEI/playlist/internal/converter"
	desc "github.com/BelyaevEI/playlist/pkg/auth_v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Login(ctx context.Context, req *desc.LoginRequest) (*desc.Response, error) {

	token, err := i.authService.Login(ctx, converter.ToLoginFromDesc(req))
	if err != nil {
		return &desc.Response{RefreshToken: ""}, status.Errorf(codes.PermissionDenied, "%s", err.Error())
	}

	return &desc.Response{RefreshToken: token}, nil
}
