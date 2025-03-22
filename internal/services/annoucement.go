package services

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type AnnouncementSvcImpl struct {
	sheetRepo SheetRepo
}

func NewAnnouncementSvcImpl(
	sheetRepo SheetRepo,
) *AnnouncementSvcImpl {
	return &AnnouncementSvcImpl{
		sheetRepo: sheetRepo,
	}
}

func (s *AnnouncementSvcImpl) SaveAnnouncement(ctx context.Context, incomingUpdate tgbotapi.Update) error {
	return s.sheetRepo.SaveMessageToAnnoucementSheet(ctx, incomingUpdate.Message)
}
