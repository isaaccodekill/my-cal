package config

import (
	"encoding/base64"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"log"
)

func CreateCalConfigContext(appConfig *Config) *oauth2.Config {
	credentialsBase64 := appConfig.CalendarCredentialsBase64
	scope := appConfig.Scope

	if credentialsBase64 == "" {
		log.Fatal("Missing GOOGLE_APPLICATION_CREDENTIALS_BASE64 environment variable")
	}

	// decode credentials from base64
	credentials, err := base64.StdEncoding.DecodeString(credentialsBase64)
	if err != nil {
		log.Fatal("Error decoding base64 credentials")
	}

	if scope == "" {
		log.Fatal("Missing GOOGLE_CALENDAR_SCOPE environment variable")
	}

	// create oauth2 config
	config, err := google.ConfigFromJSON(credentials, scope)

	if err != nil {
		log.Fatal("Error creating oauth2 config")
	}

	return config
}
