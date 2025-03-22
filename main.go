package main

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/cesc1802/english-with-me-bot/config"
	"github.com/cesc1802/english-with-me-bot/infras/sheet"
	"github.com/cesc1802/english-with-me-bot/internal/repositories"
	"github.com/cesc1802/english-with-me-bot/internal/services"
	tgbot "github.com/cesc1802/english-with-me-bot/internal/telegram_bot"
	"github.com/cesc1802/english-with-me-bot/pkg/utils"
	"github.com/joho/godotenv"
)

func main() {
	InitApp()
}

func InitApp() {
	// Initialize environment variables
	envFile := utils.LoadAdaptiveEnvFile()
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading %s file: %v", envFile, err)
	}
	fmt.Printf("Loaded environment from %s\n", envFile)

	// setup config
	var appCfg config.AppConfig
	if err := env.Parse(&appCfg); err != nil {
		log.Fatal(err)
	}

	// start dependency injections
	sheetSvc := sheet.NewSheetConn(&appCfg)
	sheetRepo := repositories.NewSheetRepoImpl(sheetSvc, appCfg.SpreadsheetId, appCfg.AnnouncementSheetName, appCfg.SubmitSheetName)
	announcementSvc := services.NewAnnouncementSvcImpl(sheetRepo)
	topicSvc := services.NewTopicSvcImpl(sheetRepo)
	// end dependency injections

	tgBot := tgbot.NewTGBot(&appCfg, topicSvc, announcementSvc)
	tgBot.Run()
}
