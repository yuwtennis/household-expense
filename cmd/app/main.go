package main

import (
	"fmt"
	"github.com/yuwtennis/household-expense/internal"
	"os"
	"time"
)

func main() {
	driveId := os.Getenv("GDRIVE_FOLDER_ID")
	bucketName := os.Getenv("BUCKET_NAME")

	// FIXME Too stateful
	time.Local, _ = time.LoadLocation("Asia/Tokyo")
	bookName := fmt.Sprintf("MonthlyBillSummary-FullTime-%d", time.Now().Year())

	internal.Run(
		driveId,
		bookName,
		bucketName)
}
