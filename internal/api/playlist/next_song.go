package playlist

import (
	"context"

	desc "github.com/BelyaevEI/playlist/pkg/playlist_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) NextSong(ctx context.Context, req *desc.Request) (*emptypb.Empty, error) {
	i.playlistService.NextSong(ctx, req.GetLogin())

	return &emptypb.Empty{}, nil
}
