package io

import (
	"cloud.google.com/go/bigquery"
	"context"
)

func NewBigQuery(projectID string) *bigquery.Client {
	ctx := context.Background()
	c, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		panic(err)
	}

	return c
}

func Write(
	client *bigquery.Client) string {
	// TODO MERGE https://cloud.google.com/bigquery/docs/reference/standard-sql/dml-syntax#merge_statement
	return "Read"
}
