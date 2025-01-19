package authtest

import (
	"context"
	"testing"

	api "github.com/BelyaevEI/playlist/internal/api/auth"
	"github.com/BelyaevEI/playlist/internal/model"
	"github.com/BelyaevEI/playlist/internal/service/auth"
	"github.com/BelyaevEI/playlist/internal/service/auth/mocks"
	desc "github.com/BelyaevEI/playlist/pkg/auth_v1"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestLogin(t *testing.T) {

	t.Parallel()
	type authServiceMockFunc func(mc *minimock.Controller) auth.AuthService

	type args struct {
		ctx context.Context
		req *desc.LoginRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		login    = gofakeit.Animal()
		password = gofakeit.Password(true, true, true, true, true, 10)
		token    = gofakeit.LoremIpsumWord()
		// serviceErr = fmt.Errorf("service error")

		req = &desc.LoginRequest{
			Login:    login,
			Password: password,
		}

		res = &desc.Response{
			RefreshToken: token,
		}

		userLogin = model.UserLogin{
			Login:    login,
			Password: password,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *desc.Response
		err             error
		authServiceMock authServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			authServiceMock: func(mc *minimock.Controller) auth.AuthService {
				mock := mocks.NewAuthServiceMock(mc)
				mock.LoginMock.Expect(ctx, &userLogin).Return(token, nil)
				return mock
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			authServiceMock := test.authServiceMock(mc)
			api := api.NewImplementation(authServiceMock)

			res, err := api.Login(test.args.ctx, test.args.req)
			require.Equal(t, test.err, err)
			require.Equal(t, test.want, res)
		})
	}

}
