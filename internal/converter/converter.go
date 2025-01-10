package converter

import (
	"time"

	"github.com/BelyaevEI/playlist/internal/model"
	desc "github.com/BelyaevEI/playlist/pkg/auth_v1"
	descPlaylist "github.com/BelyaevEI/playlist/pkg/playlist_v1"
)

func ToLoginFromDesc(user *desc.LoginRequest) *model.UserLogin {
	return &model.UserLogin{
		Login:    user.GetLogin(),
		Password: user.GetPassword(),
	}
}

func ToRegistrationFromDesc(user *desc.RegistrationRequest) *model.UserRegistration {
	return &model.UserRegistration{
		Login:           user.GetLogin(),
		Password:        user.GetPassword(),
		ConfirmPassword: user.GetConfirmPassword(),
	}
}

func ToAddSongFromDesc(song *descPlaylist.AddSongRequest) *model.SongRequest {
	return &model.SongRequest{
		Title:    song.GetTitle(),
		Article:  song.GetArticle(),
		Duration: time.Duration(song.GetDuration()),
		Login:    song.GetLogin(),
	}
}
