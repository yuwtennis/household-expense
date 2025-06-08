package services

import (
	"context"
	"fmt"
	"github.com/yuwtennis/household-expense/internal/helpers"
	"google.golang.org/api/drive/v3"
)

type GDrive struct {
	client *drive.Service
}

func NewDrive() *GDrive {
	ctx := context.Background()

	srv, err := drive.NewService(ctx)

	helpers.EvaluateErr(err, "")
	return &GDrive{client: srv}
}

func (g *GDrive) ListFilesBy(
	folderId string) []*drive.File {
	// TODO Error handling
	var files []*drive.File
	var nextPageToken string
	qString := fmt.Sprintf("'%s' in parents", folderId)

	driveListCall := g.client.Files.
		List().
		PageSize(1).
		Fields("nextPageToken, files(id, name)").
		Q(qString)

	fileList, err := driveListCall.Do()

	helpers.EvaluateErr(err, "Something went wrong while accessing google drive.")

	if len(fileList.Files) > 0 {
		files = append(files, fileList.Files...)
	}

	nextPageToken = fileList.NextPageToken

	for len(nextPageToken) > 0 {
		fileList, err = driveListCall.
			PageToken(nextPageToken).
			Do()
		helpers.EvaluateErr(err, "Something went wrong while accessing google drive.")
		nextPageToken = fileList.NextPageToken
		files = append(files, fileList.Files...)
	}

	return files
}
