package tgbot

import (
	"context"
	"log"

	"github.com/cesc1802/english-with-me-bot/config"
	"github.com/cesc1802/english-with-me-bot/internal/handlers"
	"github.com/cesc1802/english-with-me-bot/internal/services"
	"github.com/cesc1802/english-with-me-bot/pkg/statics"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TGBot struct {
	appCfg          *config.AppConfig
	topicSvc        *services.TopicSvcImpl
	announcementSvc *services.AnnouncementSvcImpl
}

func NewTGBot(
	appCfg *config.AppConfig,
	topicSvc *services.TopicSvcImpl,
	announcementSvc *services.AnnouncementSvcImpl,
) *TGBot {
	return &TGBot{
		appCfg:          appCfg,
		topicSvc:        topicSvc,
		announcementSvc: announcementSvc,
	}
}

func (b *TGBot) Run() {
	ctx := context.Background()

	// init bot
	bot, err := tgbotapi.NewBotAPI(b.appCfg.BotToken)
	if err != nil {
		log.Fatalf("Error starting Telegram bot: %v", err)
	}
	// log bot info if is in dev env
	if b.appCfg.ServiceEnv == statics.SERVICE_ENV_DEV {
		bot.Debug = true
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Set up the update listener
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// handle incoming
	incomingUpdates := bot.GetUpdatesChan(u)
	handler := handlers.NewIncomingUpdateHandler(b.appCfg, bot, incomingUpdates, b.topicSvc, b.announcementSvc)
	handler.HandleIncomingUpdates(ctx)
}
