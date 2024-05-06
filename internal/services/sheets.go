package services

import (
	"context"
	"google.golang.org/api/sheets/v4"
)

func NewSpreadSheet() *sheets.Service {
	ctx := context.Background()
	sheetService, err := sheets.NewService(ctx)

	if err != nil {
		panic(err)
	}
	return sheetService
}
