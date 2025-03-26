package handlers

import (
	"context"
	"fmt"
	"log"

	"github.com/cesc1802/english-with-me-bot/pkg/statics"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *IncomingUpdateHandler) handleTopic(ctx context.Context, incomingUpdate tgbotapi.Update) {
	err := h.topicSvc.SaveSubmit(ctx, incomingUpdate)
	if err != nil {
		log.Printf("[ERROR] Topic service error: %v", err)
	}

	// Respond to the message
	reply := tgbotapi.NewMessage(incomingUpdate.Message.Chat.ID, fmt.Sprintf("%s", statics.GetRandomMotivationalSentence(statics.MotivationalSentences)))
	reply.ReplyToMessageID = h.appCfg.StudentPresentationsTopicId
	_, err = h.bot.Send(reply)
	if err != nil {
		log.Printf("[ERROR] bot send error: %v", err)
	}
}
