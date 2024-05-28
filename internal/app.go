package internal

import (
	"context"
	"fmt"
	"github.com/yuwtennis/household-expense/internal/services"
	"google.golang.org/genproto/googleapis/type/month"
	"strings"
	"time"
)

const (
	BqExpensesTblName = "household"
)

// Run orchestrates the business logic
func Run(folderId string, bookName string, bucketName string) {
	var spreadSheetId string
	ch := make(chan *MonthlyAccount)
	ctx := context.Background()

	// Read household account book from Google Spread Sheet
	driveSrv := services.NewDrive()
	files := driveSrv.ListFilesBy(folderId)
	time.Local, _ = time.LoadLocation("Asia/Tokyo")

	for _, file := range files {
		if file.Name == bookName {
			spreadSheetId = file.Id
		}
	}

	spreadSheetSrv := services.NewSpreadSheet()

	go deserialize(spreadSheetSrv, spreadSheetId, ch)

	gcsSrv := services.NewGoogleStorage()

	for v := range ch {
		gcsSrv.Write(
			bucketName,
			fmt.Sprintf("%s/%s",
				BqExpensesTblName,
				v.AsHivePartitionLayout(),
			),
			v.Serialize(),
			ctx)
	}
}

func deserialize(
	srv *services.SpreadSheet,
	spreadSheetId string,
	ch chan *MonthlyAccount) {
	for mon := 1; mon <= 12; mon++ {
		ch <- NewMP(srv.Read(spreadSheetId, AsMonStr(mon), PaymentBookSheetRange))
	}

	close(ch)
}

func AsMonStr(mon int) string {
	monInt32 := int32(mon)
	return (month.Month_name[monInt32])[:1] + strings.ToLower(month.Month_name[monInt32])[1:3]
}
