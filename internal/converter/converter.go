package converter

import (
	"github.com/BelyaevEI/playlist/internal/model"
	desc "github.com/BelyaevEI/playlist/pkg/auth_v1"
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
