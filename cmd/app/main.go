package main

import "github.com/yuwtennis/household-expense/internal"

const (
	PaymentRecordPreix      = "MonthlyBillSummary-FullTime"
	IncomeStatementFolderId = "1nWiE2r6IUZwZSvOW8pCoyUwjVJQEJoUh"
)

func main() {
	internal.Run(IncomeStatementFolderId, PaymentRecordPreix)
}
