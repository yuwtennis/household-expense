package services

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/yuwtennis/household-expense/internal"
)

func NewGoogleStorage() *storage.Client {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	internal.EvaluateErr(err, "")

	return c
}
