package internal

import (
	"context"
	"fmt"
	"github.com/yuwtennis/household-expense/internal/services"
	"golang.org/x/oauth2/google"
	"google.golang.org/genproto/googleapis/type/month"
	"strconv"
	"strings"
	"time"
)

const (
	PaymentBookSheetRange = "B1:F57"
)

// Run orchestrates the business logic
func Run(
	gDriveFolderId string,
	accountBookPrefix string) {
	var spreadSheetId string
	var mon int32
	var book *MonthlyAccount
	ctx := context.Background()

	// Read household account book from Google Spread Sheet
	driveSrv := services.NewDrive()
	files := services.ListFilesBy(driveSrv, gDriveFolderId)

	time.Local, _ = time.LoadLocation("Asia/Tokyo")
	now := time.Now()

	thisYear := strconv.Itoa(now.Year())
	isName := accountBookPrefix + "-" + thisYear

	for _, file := range files {
		if file.Name == isName {
			spreadSheetId = file.Id
		}
	}

	spreadSheetSrv := services.NewSpreadSheet()

	// Deserialize and write to Google Storage
	for mon = 1; mon <= 12; mon++ {
		sheetName := (month.Month_name[mon])[:1] + strings.ToLower(month.Month_name[mon])[1:3]
		data, err := spreadSheetSrv.Spreadsheets.Values.Get(
			spreadSheetId,
			fmt.Sprintf("%s!%s", sheetName, PaymentBookSheetRange),
		).Do()
		EvaluateErr(err, "Error while getting data from Google Sheet.")
		book = NewMP(data.Values)

		storageSrv := services.NewGoogleStorage()
		_, err = storageSrv.
			Bucket(getCurrProjectId(ctx)).
			Object(
				"household-expense/date=" + now.Format("YYYY-MM-DD") + "book.json").
			NewWriter(ctx).
			Write(book.AsAccountRecords())

		EvaluateErr(err, "Something wrong when writing to Google Storage.")
	}
}

func getCurrProjectId(ctx context.Context) string {
	cred, err := google.FindDefaultCredentials(ctx)
	EvaluateErr(err, "")

	return cred.ProjectID
}
