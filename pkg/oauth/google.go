package oauth

import (
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GetGoogleOauthConfig() oauth2.Config {
	return oauth2.Config{
		ClientID:     viper.GetString("google_oauth.client_id"),
		ClientSecret: viper.GetString("google_oauth.client_secret"),
		RedirectURL:  viper.GetString("google_oauth.redirect_uri"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
			"email",
			"profile",
			"openid",
		},
		Endpoint: google.Endpoint,
	}
}
