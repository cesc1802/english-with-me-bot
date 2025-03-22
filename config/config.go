package config

type AppConfig struct {
	BotToken                    string `env:"BOT_TOKEN"`
	ExamGroupId                 string `env:"EXAM_GROUP_ID"`
	AnnouncementsTopicId        int    `env:"ANNOUNCEMENTS_TOPIC_ID"`
	StudentPresentationsTopicId int    `env:"STUDENT_PRESENTATIONS_TOPIC_ID"`
	QuestionAndAnswerTopicId    int    `env:"QUESTION_AND_ANSWER_TOPIC_ID"`
	SpreadsheetId               string `env:"SPREADSHEET_ID"`
	GoogleSheetCredsBase64      string `env:"GOOGLE_SHEET_CREDS_BASE64"`
	ServiceEnv                  string `env:"SERVICE_ENV"`
	AnnouncementSheetName       string `env:"ANNOUNCEMENT_SHEET_NAME"`
	SubmitSheetName             string `env:"SUBMIT_SHEET_NAME"`
}
