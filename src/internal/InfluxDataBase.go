package internal

import (
	b64 "encoding/base64"
	"fmt"
	"iot-data-viewer-backend/src/model"
	"log"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func Influx_ConfigureAndConnect(clientName string) bool {
	client := influxdb2.NewClient("http://influxdb:8086", "my-super-secret-auth-token")
	//userName := "admin"
	//password := "adminpass"
	//client := influxdb2.NewClient("http://influxdb:8086", fmt.Sprintf("%s:%s", userName, password))

	if client == nil {
		log.Println("Unable to create database client for client " + clientName)
		return false
	}

	model.MessageChannel = make(chan model.Payload)
	//writeAPI := client.WriteAPIBlocking("my-org", clientName)
	writeAPI := client.WriteAPI("my-org", "my-bucket")

	go func() {
		for {
			msgData := <-model.MessageChannel
			log.Println("Received topic:" + msgData.Topic + ", data:" + msgData.Data + " stored into the database.")
			//writeAPI.WriteRecord(fmt.Sprintf("stat data=\"%s\"", data))
			writeAPI.WriteRecord(fmt.Sprintf("stat,topic=\"%s\" data=\"%s\"", msgData.Topic, b64.StdEncoding.EncodeToString([]byte(msgData.Data))))
			//writeAPI.WriteRecord(fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 23.5, 45.0))
			// Flush writes
			writeAPI.Flush()

			// //Read data for validation
			// queryAPI := client.QueryAPI("my-org")

			// // Get parser flux query result
			// result, err := queryAPI.Query(context.Background(), `from(bucket:"my-bucket")|> range(start: -1h) |> filter(fn: (r) => r._measurement == "stat")`)
			// if err == nil {
			// 	// Use Next() to iterate over query result lines
			// 	for result.Next() {
			// 		// Observe when there is new grouping key producing new table
			// 		if result.TableChanged() {
			// 			fmt.Printf("table: %s\n", result.TableMetadata().String())
			// 		}
			// 		// read result
			// 		fmt.Printf("row: %s\n", result.Record().String())
			// 	}
			// 	if result.Err() != nil {
			// 		fmt.Printf("Query error: %s\n", result.Err().Error())
			// 	}
			// }
			// // Ensures background processes finishes
			// //client.Close()
		}
	}()

	return true
}
