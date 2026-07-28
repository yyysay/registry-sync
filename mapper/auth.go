package mapper

import (
	"registry-sync/config"
	"registry-sync/model"
)

func ConvertAuth(auth *config.AuthConfig) *model.Auth {

	if auth == nil {
		return nil
	}

	return &model.Auth{
		Username: auth.Username,
		Password: auth.Password,
	}
}
