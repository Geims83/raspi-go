package main

import (
	"fmt"
	"log"
	"net/http"
)

func read(w http.ResponseWriter, req *http.Request) {

	status, err := readData()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Fprintf(w, "Temperatura: %v\n", status.Temperature)
		fmt.Fprintf(w, "Umidità: %v\n", status.Humidity)
	}
}

func main() {

	http.HandleFunc("/", read)

	http.ListenAndServe(":8090", nil)
}
