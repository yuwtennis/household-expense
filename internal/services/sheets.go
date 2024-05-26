package services

import (
	"context"
	"github.com/yuwtennis/household-expense/internal/helpers"
	"google.golang.org/api/sheets/v4"
)

type SpreadSheet struct {
	client *sheets.Service
}

func NewSpreadSheet() *SpreadSheet {
	ctx := context.Background()
	srv, err := sheets.NewService(ctx)

	helpers.EvaluateErr(err, "")
	return &SpreadSheet{client: srv}
}

func (s *SpreadSheet) Read(
	spreadSheetId string, sheetName string, area string) [][]interface{} {
	data, err := s.client.Spreadsheets.Values.Get(
		spreadSheetId,
		sheetName+"!"+area,
	).Do()
	helpers.EvaluateErr(err, "Error while getting data from Google Sheet.")

	return data.Values
}
