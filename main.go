package main

import (
	"fmt"
	"log"
	// "net/http"
	"time"

	"context"
	"os"
	"os/signal"
    "syscall"
	"github.com/amenzhinsky/iothub/iotdevice"
	iotmqtt "github.com/amenzhinsky/iothub/iotdevice/transport/mqtt"
)

// func read(w http.ResponseWriter, req *http.Request) {

// 	status, err := readData()
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		fmt.Fprintf(w, "Temperatura: %v\n", status.Temperature)
// 		fmt.Fprintf(w, "Umidità: %v\n", status.Humidity)
// 	}
// }

// func main_web() {

// 	http.HandleFunc("/", read)

// 	http.ListenAndServe(":8090", nil)
// }

func main() {

	client, err := iotdevice.NewFromConnectionString(
		iotmqtt.New(), os.Getenv("IOTHUB_DEVICE_CONNECTION_STRING"),
	)
	if err != nil {
		log.Fatal(err)
	}
	
	// connect to the iothub
	if err = client.Connect(context.Background()); err != nil {
		log.Fatal(err)
	}

	ticker := time.NewTicker(5 * time.Second)
    done := make(chan bool)
    
	go func() {
        for {
            select {
            case <-done:
                return
            case <-ticker.C:
				status, err := readData()

				if err != nil {
					log.Fatal(err)
				} else {
					
					// if err = h.SendEvent(context.Background(), []byte(`hello`)); err != nil {
					// 	log.Fatal(err)
					// }
					fmt.Printf("Temperatura: %v\n", status.Temperature)
					fmt.Printf("Umidità: %v\n", status.Humidity)
				}
            }
        }
    }()

	c := make(chan os.Signal)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
		ticker.Stop()
		done <- true
		fmt.Println("Ticker stopped")
    }()

}
