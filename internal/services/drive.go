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

func NewDrive() (*GDrive, *helpers.AppErr) {
	ctx := context.Background()

	srv, err := drive.NewService(ctx)

	if err != nil {
		return nil, &helpers.AppErr{
			Error: err,
			Msg:   "Error creating drive client",
		}
	}

	return &GDrive{client: srv}, nil
}

func (g *GDrive) ListFilesBy(
	folderId string) ([]*drive.File, *helpers.AppErr) {
	var files []*drive.File
	var nextPageToken string
	qString := fmt.Sprintf("'%s' in parents", folderId)

	driveListCall := g.client.Files.
		List().
		PageSize(1).
		Fields("nextPageToken, files(id, name)").
		Q(qString)

	fileList, dListErr := driveListCall.Do()

	if dListErr != nil {
		return nil, &helpers.AppErr{
			Error: dListErr,
			Msg:   "Something went wrong while accessing google drive.",
		}
	}

	if len(fileList.Files) > 0 {
		files = append(files, fileList.Files...)
	}

	nextPageToken = fileList.NextPageToken

	for len(nextPageToken) > 0 {
		fileList, dListErr = driveListCall.
			PageToken(nextPageToken).
			Do()
		if dListErr != nil {
			return nil, &helpers.AppErr{
				Error: dListErr,
				Msg:   "Something went wrong while accessing google drive.",
			}
		}
		nextPageToken = fileList.NextPageToken
		files = append(files, fileList.Files...)
	}

	return files, nil
}
