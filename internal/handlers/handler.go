package handlers

import (
	"context"
	"log"
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

type MemberSvc interface {
	SaveMember(ctx context.Context, incomingUpdate tgbotapi.Update) error
}

type IncomingUpdateHandler struct {
	appCfg           *config.AppConfig
	bot              *tgbotapi.BotAPI
	incommingUpdates tgbotapi.UpdatesChannel
	topicSvc         TopicSvc
	announcementSvc  AnnouncementSvc
	memberSvc        MemberSvc
}

func NewIncomingUpdateHandler(
	appCfg *config.AppConfig,
	bot *tgbotapi.BotAPI,
	incommingUpdates tgbotapi.UpdatesChannel,
	topicSvc TopicSvc,
	announcementSvc AnnouncementSvc,
	memberSvc MemberSvc,
) *IncomingUpdateHandler {
	return &IncomingUpdateHandler{
		appCfg:           appCfg,
		bot:              bot,
		incommingUpdates: incommingUpdates,
		topicSvc:         topicSvc,
		announcementSvc:  announcementSvc,
		memberSvc:        memberSvc,
	}
}

func (h *IncomingUpdateHandler) HandleIncomingUpdates(ctx context.Context) {
	// Process incoming messages
	for incomingUpdate := range h.incommingUpdates {
		if incomingUpdate.Message != nil && len(incomingUpdate.Message.NewChatMembers) > 0 {
			h.handleNewChatMembers(ctx, incomingUpdate)
			continue
		}

		// handle topic
		if incomingUpdate.Message != nil && (strings.Contains(incomingUpdate.Message.Text, "#submit") || strings.Contains(incomingUpdate.Message.Caption, "#submit")) {
			h.handleTopic(ctx, incomingUpdate)
			continue
		}

		// handle annoucement
		if incomingUpdate.Message != nil && strings.Contains(incomingUpdate.Message.Text, "#topic") {
			h.handleAnnouncement(ctx, incomingUpdate)
			continue
		}

		// handle download photo command
		// Check if the message is a command
		if incomingUpdate.Message.IsCommand() {
			// Prepare a response message
			msg := tgbotapi.NewMessage(incomingUpdate.Message.Chat.ID, "")
			args := incomingUpdate.Message.CommandArguments()

			// Handle different commands
			switch incomingUpdate.Message.Command() {
			case "get_photo_id":
				msg.Text = GetPhotoURL(h.bot, strings.TrimSpace(args))
			default:
				msg.Text = "I don't know that command"
			}

			// Send the response
			if _, err := h.bot.Send(msg); err != nil {
				log.Panic(err)
			}

			continue
		}
	}
}

func GetPhotoURL(bot *tgbotapi.BotAPI, fileID string) string {
	url, _ := bot.GetFileDirectURL(fileID)
	return url
}
