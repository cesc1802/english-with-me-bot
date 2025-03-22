package handlers

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *IncomingUpdateHandler) handleAnnouncement(ctx context.Context, incomingUpdate tgbotapi.Update) {
	err := h.announcementSvc.SaveAnnouncement(ctx, incomingUpdate)
	if err != nil {
		log.Printf("[ERROR] Topic service error: %v", err)
	}

	// Respond to the message
	reply := tgbotapi.NewMessage(incomingUpdate.Message.Chat.ID, "Submitted!")
	reply.ReplyToMessageID = h.appCfg.StudentPresentationsTopicId
	_, err = h.bot.Send(reply)
	if err != nil {
		log.Printf("[ERROR] bot send error: %v", err)
	}
}
