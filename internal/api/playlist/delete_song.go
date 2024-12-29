package playlist

import (
	"context"

	"github.com/BelyaevEI/playlist/internal/converter"
	desc "github.com/BelyaevEI/playlist/pkg/playlist_v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) DeleteSong(ctx context.Context, req *desc.AddSongRequest) (*emptypb.Empty, error) {
	err := i.playlistService.DeleteSong(ctx, converter.ToAddSongFromDesc(req))
	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "%s", err.Error())
	}

	return &emptypb.Empty{}, nil
}
