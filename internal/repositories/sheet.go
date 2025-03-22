package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/cesc1802/english-with-me-bot/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"google.golang.org/api/sheets/v4"
)

type SheetRepoImpl struct {
	sheetSvc              *sheets.Service
	spreadsheetFileId     string
	announcementSheetName string
	submitSheetName       string
}

func NewSheetRepoImpl(
	sheetSvc *sheets.Service,
	spreadsheetFileId string,
	announcementSheetName string,
	submitSheetName string,
) *SheetRepoImpl {
	return &SheetRepoImpl{
		sheetSvc:              sheetSvc,
		spreadsheetFileId:     spreadsheetFileId,
		announcementSheetName: announcementSheetName,
		submitSheetName:       submitSheetName,
	}
}

func (r *SheetRepoImpl) SaveMessageToAnnoucementSheet(ctx context.Context, tgBotAPIMessage *tgbotapi.Message) error {
	timestamp := time.Now().Format(time.RFC3339)
	username := tgBotAPIMessage.From.UserName
	fullname := fmt.Sprintf("%s %s", tgBotAPIMessage.From.FirstName, tgBotAPIMessage.From.LastName)
	text := tgBotAPIMessage.Text

	annoucement := models.AnnoucementSheet{
		Day:                timestamp,
		Time:               timestamp,
		Username:           username,
		Fullname:           fullname,
		AnnoucementContent: text,
	}

	values := [][]interface{}{
		annoucement.ToSheetValue(),
	}

	valueRange := &sheets.ValueRange{
		Values: values,
	}

	_, err := r.sheetSvc.Spreadsheets.Values.Append(
		r.spreadsheetFileId,
		fmt.Sprintf("%s!A1:E1", r.announcementSheetName),
		valueRange,
	).ValueInputOption("USER_ENTERED").InsertDataOption("INSERT_ROWS").Do()

	if err != nil {
		return err
	}

	return nil
}

func (r *SheetRepoImpl) SaveMessageToSubmitSheet(ctx context.Context, tgBotAPIMessage *tgbotapi.Message) error {
	timestamp := time.Now().Format(time.RFC3339)
	username := tgBotAPIMessage.From.UserName
	fullname := fmt.Sprintf("%s %s", tgBotAPIMessage.From.FirstName, tgBotAPIMessage.From.LastName)
	text := tgBotAPIMessage.Text

	submitSheet := models.SubmitSheet{
		Day:           timestamp,
		Time:          timestamp,
		Username:      username,
		Fullname:      fullname,
		SubmitContent: text,
	}

	values := [][]interface{}{
		submitSheet.ToSheetValue(),
	}

	valueRange := &sheets.ValueRange{
		Values: values,
	}

	_, err := r.sheetSvc.Spreadsheets.Values.Append(
		r.spreadsheetFileId,
		fmt.Sprintf("%s!A1:E1", r.submitSheetName),
		valueRange,
	).ValueInputOption("USER_ENTERED").InsertDataOption("INSERT_ROWS").Do()

	if err != nil {
		return err
	}

	return nil
}
