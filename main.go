package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func main() {
	token := os.Getenv("INFLUXDB_TOKEN")
	url := "http://localhost:8086"
	client := influxdb2.NewClient(url, token)

	// Write some data
	org := "pomatti"
	bucket := "temp"
	writeAPI := client.WriteAPIBlocking(org, bucket)
	for value := 0; value < 5; value++ {
		tags := map[string]string{
			"unit": "temperature",
		}
		fields := map[string]interface{}{
			"avg": 24.5,
			"max": 45,
		}
		p := influxdb2.NewPoint("stat", tags, fields, time.Now())
		time.Sleep(1 * time.Second) // separate points by 1 second

		if err := writeAPI.WritePoint(context.Background(), p); err != nil {
			log.Fatal(err)
		}
	}

	// Simple query
	queryAPI := client.QueryAPI(org)
	query := `from(bucket: "temp")
            |> range(start: -10m)
            |> filter(fn: (r) => r._measurement == "stat")`
	results, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	for results.Next() {
		fmt.Println(results.Record())
	}
	if err := results.Err(); err != nil {
		log.Fatal(err)
	}

	// Aggregate query
	query = `from(bucket: "temp")
              |> range(start: -10m)
              |> filter(fn: (r) => r._measurement == "stat")
              |> mean()`
	results, err = queryAPI.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	for results.Next() {
		fmt.Println(results.Record())
	}
	if err := results.Err(); err != nil {
		log.Fatal(err)
	}
}
