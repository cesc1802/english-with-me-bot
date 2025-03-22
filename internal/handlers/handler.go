package handlers

import (
	"context"
	"strings"

	"github.com/cesc1802/english-with-me-bot/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TopicSvc interface {
	SaveSubmit(ctx context.Context, incomingUpdate tgbotapi.Update) error
}

type AnnouncementSvc interface {
	SaveAnnouncement(ctx context.Context, incomingUpdate tgbotapi.Update) error
}

type IncomingUpdateHandler struct {
	appCfg           *config.AppConfig
	bot              *tgbotapi.BotAPI
	incommingUpdates tgbotapi.UpdatesChannel
	topicSvc         TopicSvc
	announcementSvc  AnnouncementSvc
}

func NewIncomingUpdateHandler(
	appCfg *config.AppConfig,
	bot *tgbotapi.BotAPI,
	incommingUpdates tgbotapi.UpdatesChannel,
	topicSvc TopicSvc,
	announcementSvc AnnouncementSvc,
) *IncomingUpdateHandler {
	return &IncomingUpdateHandler{
		appCfg:           appCfg,
		bot:              bot,
		incommingUpdates: incommingUpdates,
		topicSvc:         topicSvc,
		announcementSvc:  announcementSvc,
	}
}

func (h *IncomingUpdateHandler) HandleIncomingUpdates(ctx context.Context) {
	// Process incoming messages
	for incomingUpdate := range h.incommingUpdates {
		// handle topic
		if incomingUpdate.Message != nil && strings.Contains(incomingUpdate.Message.Text, "#topic") {
			h.handleTopic(ctx, incomingUpdate)
		}

		// handle annoucement
		if incomingUpdate.Message != nil && strings.Contains(incomingUpdate.Message.Text, "#submit") {
			h.handleAnnouncement(ctx, incomingUpdate)
		}
	}
}
