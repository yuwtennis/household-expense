package services

import (
	"context"
	"github.com/yuwtennis/household-expense/internal/helpers"
	"google.golang.org/api/sheets/v4"
)

type SpreadSheet struct {
	client *sheets.Service
}

func NewSpreadSheet() (*SpreadSheet, *helpers.AppErr) {
	ctx := context.Background()
	srv, err := sheets.NewService(ctx)

	if err != nil {
		return nil, &helpers.AppErr{
			Error: err,
			Msg:   "Error creating sheet client",
		}
	}

	return &SpreadSheet{client: srv}, nil
}

func (s *SpreadSheet) Read(
	spreadSheetId string, sheetName string, area string) ([][]interface{}, *helpers.AppErr) {
	data, err := s.client.Spreadsheets.Values.Get(
		spreadSheetId,
		sheetName+"!"+area,
	).Do()

	if err != nil {
		return nil, &helpers.AppErr{
			Error: err,
			Msg:   "Error while getting data from Google Sheet.",
		}
	}

	return data.Values, nil
}
