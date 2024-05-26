package main

import "github.com/yuwtennis/household-expense/internal"

const (
	PaymentRecordBookName   = "MonthlyBillSummary-FullTime-2024"
	IncomeStatementFolderId = "1nWiE2r6IUZwZSvOW8pCoyUwjVJQEJoUh"
	BucketName              = "elite-caster-125113"
)

func main() {
	internal.Run(
		IncomeStatementFolderId,
		PaymentRecordBookName,
		BucketName)
}
