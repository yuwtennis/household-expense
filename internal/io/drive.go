package io

import (
	"context"
	"google.golang.org/api/drive/v3"
)

func NewGDrive() *drive.Service {
	ctx := context.Background()
	driveServe, err := drive.NewService(ctx)

	if err != nil {
		panic(err)
	}
	return driveServe
}
