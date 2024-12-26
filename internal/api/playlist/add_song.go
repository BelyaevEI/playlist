package playlist

import (
	"context"

	"github.com/BelyaevEI/playlist/internal/converter"
	desc "github.com/BelyaevEI/playlist/pkg/playlist_v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) AddSong(ctx context.Context, req *desc.AddSongRequest) (*desc.Response, error) {

	err := i.playlistService.AddSong(ctx, converter.ToAddSongFromDesc(req))
	if err != nil {
		return &desc.Response{}, status.Errorf(codes.Internal, "%s", err.Error())
	}

	return &desc.Response{RefreshToken: ""}, nil
}
