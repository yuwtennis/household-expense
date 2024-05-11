package internal

import (
	"fmt"
	"github.com/yuwtennis/household-expense/internal/services"
	"google.golang.org/genproto/googleapis/type/month"
	"strconv"
	"strings"
	"time"
)

const (
	PaymentBookSheetRange = "B1:F57"
)

// Run orchestrates the business logic
func Run(folderId string, prefix string) {
	var spreadSheetId string
	var m int32
	var book *MonthlyPayment

	driveSrv := services.NewDrive()
	files := services.ListFilesBy(driveSrv, folderId)

	time.Local, _ = time.LoadLocation("Asia/Tokyo")

	thisYear := strconv.Itoa(time.Now().Year())
	isName := prefix + "-" + thisYear

	for _, file := range files {
		if file.Name == isName {
			spreadSheetId = file.Id
		}
	}

	spreadSheetSrv := services.NewSpreadSheet()

	for m = 1; m <= 12; m++ {
		sheetName := (month.Month_name[m])[:1] + strings.ToLower(month.Month_name[m])[1:]
		data, err := spreadSheetSrv.Spreadsheets.Values.Get(
			spreadSheetId,
			fmt.Sprintf("%s!%s", sheetName, PaymentBookSheetRange),
		).Do()
		EvaluateErr(err, "Error while getting data from Sheet.")
		book = NewMP(data.Values)
	}

	// TODO Merge to bigquery
	bqClient := services.NewBigQuery("MyProjectId")
	services.Write(bqClient, book)
}
