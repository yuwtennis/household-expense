package services

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/yuwtennis/household-expense/internal/helpers"
)

type GCS struct {
	client *storage.Client
}

func NewGoogleStorage() *GCS {
	ctx := context.Background()
	srv, err := storage.NewClient(ctx)
	helpers.EvaluateErr(err, "Something went wrong when creating a client.")

	return &GCS{client: srv}
}

func (g *GCS) Write(
	bucketName string,
	folderPath string,
	b []byte,
	ctx context.Context) {

	wc := g.client.
		Bucket(bucketName).
		Object(
			folderPath).
		NewWriter(ctx)
	wc.ContentType = "text/plain"
	_, errW := wc.Write(b)
	helpers.EvaluateErr(errW, "Something wrong when writing to Google Storage.")

	errC := wc.Close()
	helpers.EvaluateErr(errC, "Something wrong when closing writer.")
}
