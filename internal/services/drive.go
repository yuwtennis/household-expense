package services

import (
	"context"
	"fmt"
	"google.golang.org/api/drive/v3"
)

func NewDrive() *drive.Service {
	ctx := context.Background()
	driveServe, err := drive.NewService(ctx)

	if err != nil {
		panic(err)
	}
	return driveServe
}

func ListFilesBy(
	srv *drive.Service,
	folderId string) []*drive.File {
	// TODO Error handling
	var files []*drive.File
	var nextPageToken string
	qString := fmt.Sprintf("'%s' in parents", folderId)

	driveListCall := srv.Files.
		List().
		PageSize(1).
		Fields("nextPageToken, files(id, name)").
		Q(qString)

	fileList, _ := driveListCall.Do()

	if len(fileList.Files) > 0 {
		files = append(files, fileList.Files...)
	}

	nextPageToken = fileList.NextPageToken

	for len(nextPageToken) > 0 {
		fileList, _ = driveListCall.
			PageToken(nextPageToken).
			Do()

		nextPageToken = fileList.NextPageToken
		files = append(files, fileList.Files...)
	}

	return files
}
