package sheet

import (
	"context"
	"encoding/base64"
	"log"

	"github.com/cesc1802/english-with-me-bot/config"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func NewSheetConn(appCfg *config.AppConfig) *sheets.Service {
	// load google sheets api credentials
	credBytes, err := base64.StdEncoding.DecodeString(appCfg.GoogleSheetCredsBase64)
	if err != nil {
		log.Fatalf("Failed to decode credentials: %v", err)
	}
	// authen google
	config, err := google.JWTConfigFromJSON(credBytes, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatalf("Failed to create JWT config: %v", err)
	}
	// init sheets service
	client := config.Client(context.Background())
	srv, err := sheets.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Failed to create Sheets service: %v", err)
	}

	return srv
}
