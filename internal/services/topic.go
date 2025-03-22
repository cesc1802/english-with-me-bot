package services

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TopicSvcImpl struct {
	sheetRepo SheetRepo
}

func NewTopicSvcImpl(
	sheetRepo SheetRepo,
) *TopicSvcImpl {
	return &TopicSvcImpl{
		sheetRepo: sheetRepo,
	}
}

func (s *TopicSvcImpl) SaveSubmit(ctx context.Context, incomingUpdate tgbotapi.Update) error {
	return s.sheetRepo.SaveMessageToSubmitSheet(ctx, incomingUpdate.Message)
}
