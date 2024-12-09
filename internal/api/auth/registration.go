package auth

import (
	"context"

	"github.com/BelyaevEI/playlist/internal/converter"
	desc "github.com/BelyaevEI/playlist/pkg/auth_v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Registration(ctx context.Context, req *desc.RegistrationRequest) (*desc.Response, error) {

	token, err := i.authService.Registration(ctx, converter.ToRegistrationFromDesc(req))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &desc.Response{RefreshToken: token}, nil
}
