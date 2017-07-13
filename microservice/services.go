package microservice

import (
	"context"
	"log"
	"os"
	"strings"

	"google.golang.org/api/iterator"
	"google.golang.org/api/option"

	"cloud.google.com/go/bigquery"
)

func queryService(sql string) ([]map[string]bigquery.Value, error) {
	log.Println("[SERVICE] Doing query")
	ctx := context.Background()
	projectID := os.Getenv("GCLOUD_PROJECT_ID")
	client, err := bigquery.NewClient(ctx, projectID, option.WithServiceAccountFile("./credentials.json"))

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if strings.Contains(sql, "viz_inside") {
		sql = `CREATE TEMP FUNCTION viz_inside(long FLOAT64, lat FLOAT64, poly STRING)
			  RETURNS BOOLEAN
			  LANGUAGE js AS
			"""
			    return VizGeo.inside(long, lat, JSON.parse(poly));
			"""
			OPTIONS (
			  library="gs://bigquery-geospatial-viz/vizgeo.min.js"
			);` + sql
	}
	if strings.Contains(sql, "viz_intersect") {
		sql = `CREATE TEMP FUNCTION viz_intersect(long FLOAT64, lat FLOAT64, poly STRING)
			  RETURNS BOOLEAN
			  LANGUAGE js AS
			"""
				return VizGeo.intersect(long, lat, JSON.parse(poly));
			"""
			OPTIONS (
			  library="gs://bigquery-geospatial-viz/vizgeo.min.js"
			);` + sql
	}

	q := client.Query(sql)
	q.UseStandardSQL = true
	it, err := q.Read(ctx)

	if err != nil {
		log.Fatal(err)
		return nil, err
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
			log.Fatal(err)
			return nil, err
		}
	}

	log.Println("[SERVICE] Query finished")
	return rows, nil
}
