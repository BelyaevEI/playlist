package playlisttest

import (
	"context"
	"testing"

	"github.com/BelyaevEI/playlist/internal/logger"
	"github.com/BelyaevEI/playlist/internal/model"
	"github.com/BelyaevEI/playlist/internal/repository/playlist"
	"github.com/BelyaevEI/playlist/internal/repository/playlist/mocks"
	service "github.com/BelyaevEI/playlist/internal/service/playlist"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestAddSong(t *testing.T) {

	t.Parallel()
	type playlistRepositoryMockFunc func(mc *minimock.Controller) playlist.PlaylistRepository

	type args struct {
		ctx context.Context
		req *model.SongRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		login   = gofakeit.Animal()
		title   = gofakeit.BookTitle()
		article = gofakeit.BookAuthor()

		song = model.SongRequest{
			Login:   login,
			Article: article,
			Title:   title,
		}
	)

	tests := []struct {
		name                   string
		args                   args
		err                    error
		playlistRepositoryMock playlistRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: &song,
			},
			err: nil,
			playlistRepositoryMock: func(mc *minimock.Controller) playlist.PlaylistRepository {
				mock := mocks.NewPlaylistRepositoryMock(mc)
				logger.Init(logger.GetCore(logger.GetAtomicLevel("DEBUG")))
				mock.AddSongMock.Expect(ctx, &song).Return(nil)
				return mock
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			playlistRepositoryMock := test.playlistRepositoryMock(mc)
			service := service.NewService(playlistRepositoryMock, make(chan string))

			err := service.AddSong(test.args.ctx, test.args.req)
			require.Equal(t, test.err, err)
		})
	}

}
