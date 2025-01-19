package playlisttest

import (
	"context"
	"testing"
	"time"

	api "github.com/BelyaevEI/playlist/internal/api/playlist"
	"github.com/BelyaevEI/playlist/internal/model"
	"github.com/BelyaevEI/playlist/internal/service/playlist"
	"github.com/BelyaevEI/playlist/internal/service/playlist/mocks"
	desc "github.com/BelyaevEI/playlist/pkg/playlist_v1"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestDeleteSong(t *testing.T) {

	t.Parallel()
	type playlistServiceMockFunc func(mc *minimock.Controller) playlist.PlaylistService

	type args struct {
		ctx context.Context
		req *desc.AddSongRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		login    = gofakeit.Animal()
		title    = gofakeit.BookTitle()
		article  = gofakeit.BookAuthor()
		duration = gofakeit.Int32()
		// serviceErr = fmt.Errorf("service error")

		req = &desc.AddSongRequest{
			Login:    login,
			Title:    title,
			Article:  article,
			Duration: duration,
		}

		songReg = model.SongRequest{
			Login:    login,
			Title:    title,
			Article:  article,
			Duration: time.Duration(duration) * time.Second,
		}
	)

	tests := []struct {
		name                string
		args                args
		err                 error
		playlistServiceMock playlistServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: nil,
			playlistServiceMock: func(mc *minimock.Controller) playlist.PlaylistService {
				mock := mocks.NewPlaylistServiceMock(mc)
				mock.DeleteSongMock.Expect(ctx, &songReg).Return(nil)
				return mock
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			playlistServiceMock := test.playlistServiceMock(mc)
			api := api.NewImplementation(playlistServiceMock)

			_, err := api.DeleteSong(test.args.ctx, test.args.req)
			require.Equal(t, test.err, err)
		})
	}

}
