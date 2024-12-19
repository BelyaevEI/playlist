package playlist

import (
	"context"

	desc "github.com/BelyaevEI/playlist/pkg/playlist_v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) AddSong(ctx context.Context, req *desc.AddSongRequest) (*emptypb.Empty, error) {

	return &emptypb.Empty{}, nil
}
