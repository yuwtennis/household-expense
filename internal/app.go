package internal

import (
	"fmt"
	"github.com/yuwtennis/household-expense/internal/bop"
	"github.com/yuwtennis/household-expense/internal/services"
	"golang.org/x/exp/slices"
	"strconv"
	"time"
)

// Run orchestrates the business logic
func Run(folderId string, prefix string) {
	// Call the function from the package
	// "github.com/username/golang-package-example"
	var spreadSheetId string
	var incomes []bop.Income
	var expenses []*bop.Expense

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
	spreadSheet, _ := spreadSheetSrv.Spreadsheets.Get(spreadSheetId).Do()

	for _, sheet := range spreadSheet.Sheets {
		fmt.Printf("%s , %d\n", sheet.Properties.Title, sheet.Properties.SheetId)

		if isValidCalMonth(sheet.Properties.Title) {
			expenses = append(expenses, bop.NewExpense(sheet.Data))
			incomes = append(incomes, bop.NewMonthlySalary(sheet.Data))
		}
	}

	// TODO Merge to bigquery
}

func isValidCalMonth(monthName string) bool {
	calMonths := []string{
		"Jan",
		"Feb",
		"Mar",
		"Apr",
		"May",
		"Jun",
		"Jul",
		"Aug",
		"Sep",
		"Oct",
		"Nov",
		"Dec",
	}
	return slices.Contains(calMonths, monthName)
}
