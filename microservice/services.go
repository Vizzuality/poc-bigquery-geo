package microservice

import (
	"context"
	"log"
	"os"

	"google.golang.org/api/iterator"
	"google.golang.org/api/option"

	"cloud.google.com/go/bigquery"
)

func queryService(sql string) []map[string]bigquery.Value {
	log.Println("[SERVICE] Doing query")
	ctx := context.Background()
	projectID := os.Getenv("GCLOUD_PROJECT_ID")
	client, err := bigquery.NewClient(ctx, projectID, option.WithServiceAccountFile("./credentials.json"))

	if err != nil {
		// TODO: Handle error.
		log.Fatal(err)
	}

	q := client.Query(sql)
	it, err := q.Read(ctx)

	if err != nil {
		// TODO: Handle error.
		log.Fatal(err)
	}

	var rows []map[string]bigquery.Value
	var row map[string]bigquery.Value
	for {
		err := it.Next(&row)
		rows = append(rows, row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
	}
	log.Println("[SERVICE] Query finished")
	return rows
}
