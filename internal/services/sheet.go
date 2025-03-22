package services

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SheetRepo interface {
	SaveMessageToAnnoucementSheet(ctx context.Context, tgBotAPIMessage *tgbotapi.Message) error
	SaveMessageToSubmitSheet(ctx context.Context, tgBotAPIMessage *tgbotapi.Message) error
}
