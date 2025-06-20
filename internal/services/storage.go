package services

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/yuwtennis/household-expense/internal/helpers"
)

type GCS struct {
	client *storage.Client
}

func NewGoogleStorage() (*GCS, *helpers.AppErr) {
	ctx := context.Background()
	srv, err := storage.NewClient(ctx)

	if err != nil {
		return nil, &helpers.AppErr{
			Error: err,
			Msg:   "Something went wrong when creating a client.",
		}
	}

	return &GCS{client: srv}, nil
}

func (g *GCS) Write(
	bucketName string,
	folderPath string,
	b []byte,
	ctx context.Context) *helpers.AppErr {

	wc := g.client.
		Bucket(bucketName).
		Object(
			folderPath).
		NewWriter(ctx)

	defer func(w *storage.Writer) {
		_ = w.Close()
	}(wc)

	wc.ContentType = "text/plain"
	_, errW := wc.Write(b)
	if errW != nil {
		return &helpers.AppErr{
			Error: errW,
			Msg:   "Something wrong when writing to Google Storage.",
		}
	}

	return nil
}
