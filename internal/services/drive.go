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

func (g *GDrive) GetFile(
	folderId string,
	bookName string) (string, *helpers.AppErr) {
	qString := fmt.Sprintf("'%s' in parents and name = '%s'", folderId, bookName)

	driveListCall := g.client.Files.
		List().
		PageSize(1).
		Fields("files(id, name)").
		Q(qString)

	file, dListErr := driveListCall.Do()

	if dListErr != nil {
		return "", &helpers.AppErr{
			Error: dListErr,
			Msg:   "Something went wrong while accessing google drive.",
		}
	}

	if len(file.Files) < 1 {
		return "", &helpers.AppErr{
			Error: nil,
			Msg:   fmt.Sprintf("File does not exist. filename: %s", bookName),
		}
	}

	return file.Files[len(file.Files)-1].Id, nil
}
