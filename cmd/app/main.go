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

	time.Local, _ = time.LoadLocation("Asia/Tokyo")
	bookName := fmt.Sprintf("MonthlyBillSummary-FullTime-%d", time.Now().Year())

	err := internal.Run(
		driveId,
		bookName,
		bucketName)

	if err != nil {
		panic(fmt.Sprintf("ERROR: Runtime error. msg: %s, error: %v", err.Msg, err.Error))
	}
}
