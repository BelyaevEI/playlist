package authtest

import (
	"context"
	"testing"

	"github.com/BelyaevEI/playlist/internal/logger"
	"github.com/BelyaevEI/playlist/internal/model"
	"github.com/BelyaevEI/playlist/internal/repository/auth"
	"github.com/BelyaevEI/playlist/internal/repository/auth/mocks"
	service "github.com/BelyaevEI/playlist/internal/service/auth"
	"golang.org/x/crypto/bcrypt"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestLogin(t *testing.T) {

	t.Parallel()
	type authRepositoryMockFunc func(mc *minimock.Controller) auth.AuthRepository

	type args struct {
		ctx context.Context
		req *model.UserLogin
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		login             = gofakeit.Animal()
		password          = gofakeit.Password(true, true, true, true, true, 20)
		hashedPassword, _ = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		secretWord        = gofakeit.BeerName()
		token             = gofakeit.LoremIpsumWord()

		userLogin = model.UserLogin{
			Login:      login,
			Password:   password,
			SecretWord: secretWord,
		}

		res = model.UserLogin{
			Login:      login,
			Password:   string(hashedPassword),
			SecretWord: secretWord,
		}
	)

	tests := []struct {
		name               string
		args               args
		want               string
		err                error
		authRepositoryMock authRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: &userLogin,
			},
			want: token,
			err:  nil,
			authRepositoryMock: func(mc *minimock.Controller) auth.AuthRepository {
				mock := mocks.NewAuthRepositoryMock(mc)
				logger.Init(logger.GetCore(logger.GetAtomicLevel("DEBUG")))
				mock.SelectUserMock.Expect(ctx, &userLogin).Return(res, nil)
				return mock
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			authRepositoryMock := test.authRepositoryMock(mc)
			service := service.NewService(authRepositoryMock)

			res, err := service.Login(test.args.ctx, test.args.req)
			require.Equal(t, test.err, err)
			require.NotEmpty(t, res)
		})
	}

}
