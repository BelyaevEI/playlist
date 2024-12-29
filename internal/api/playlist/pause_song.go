package playlist

import (
	"context"

	desc "github.com/BelyaevEI/playlist/pkg/playlist_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) PauseSong(ctx context.Context, req *desc.Request) (*emptypb.Empty, error) {
	i.playlistService.PauseSong(ctx, req.GetLogin())

	return &emptypb.Empty{}, nil
}
