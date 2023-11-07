package services

import (
	"fmt"
	"golang.org/x/oauth2"
)

var Client_id string = "538172717799-agmf1rd531sol80jfgn9lgb55pe4kr4v.apps.googleusercontent.com"
var Client_secret string = "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"

// Request a token from the web, then returns the retrieved token.
func GenerateAuthUrl(config *oauth2.Config) string {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	return authURL
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
}
