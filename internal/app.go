package internal

import (
	"context"
	"fmt"
	"github.com/yuwtennis/household-expense/internal/helpers"
	"github.com/yuwtennis/household-expense/internal/services"
	"google.golang.org/genproto/googleapis/type/month"
	"strings"
	"time"
)

const (
	BqExpensesTblName = "household"
)

// Run orchestrates the business logic
func Run(folderId string, bookName string, bucketName string) *helpers.AppErr {
	var spreadSheetId string
	ch := make(chan *MonthlyAccount)
	ctx := context.Background()

	// Read household account book from Google Spread Sheet
	driveSrv, driveErr := services.NewDrive()

	if driveErr != nil {
		return driveErr
	}

	files, listErr := driveSrv.ListFilesBy(folderId)

	if listErr != nil {
		return listErr
	}

	time.Local, _ = time.LoadLocation("Asia/Tokyo")

	for _, file := range files {
		if file.Name == bookName {
			spreadSheetId = file.Id
		}
	}

	spreadSheetSrv, sheetErr := services.NewSpreadSheet()

	if sheetErr != nil {
		return sheetErr
	}

	go deserialize(spreadSheetSrv, spreadSheetId, ch)

	gcsSrv, gcsErr := services.NewGoogleStorage()

	if gcsErr != nil {
		return gcsErr
	}

	for v := range ch {
		data, serErr := v.Serialize()

		// CAPS theorum. Availability goes over consistency
		if serErr != nil {
			fmt.Printf("WARN: Failed to serialize. msg: %s,  err: %v",
				serErr.Msg, serErr.Error)
		}

		gcsSrv.Write(
			bucketName,
			fmt.Sprintf("%s/%s",
				BqExpensesTblName,
				v.AsHivePartitionLayout(),
			),
			data,
			ctx)
	}

	return nil
}

func deserialize(
	srv *services.SpreadSheet,
	spreadSheetId string,
	ch chan *MonthlyAccount) {
	for mon := 1; mon <= 12; mon++ {
		data, err := srv.Read(spreadSheetId, AsMonStr(mon), PaymentBookSheetRange)

		if err != nil {
			fmt.Printf("WARN: Failed to deser sheets. msg: %s , err: %v", err.Msg, err.Error)
		}

		ch <- NewMP(data)
	}

	close(ch)
}

func AsMonStr(mon int) string {
	monInt32 := int32(mon)
	return (month.Month_name[monInt32])[:1] + strings.ToLower(month.Month_name[monInt32])[1:3]
}
