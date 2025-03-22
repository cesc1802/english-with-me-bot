package models

type AnnoucementSheet struct {
	Day                string
	Time               string
	Username           string
	Fullname           string
	AnnoucementContent string
}

func (a AnnoucementSheet) ToSheetValue() []any {
	return []any{
		a.Day,
		a.Time,
		a.Username,
		a.Fullname,
		a.AnnoucementContent,
	}
}

type SubmitSheet struct {
	Day           string
	Time          string
	Username      string
	Fullname      string
	SubmitContent string
}

func (a SubmitSheet) ToSheetValue() []any {
	return []any{
		a.Day,
		a.Time,
		a.Username,
		a.Fullname,
		a.SubmitContent,
	}
}
