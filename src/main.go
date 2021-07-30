package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var number int

func main() {
	number = 0
	http.HandleFunc("/", display)
	http.HandleFunc("/plus", plus)
	http.HandleFunc("/minus", minus)
	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(":8181", nil))
}

func display(writer http.ResponseWriter, request *http.Request) {
	execute(writer, 0)
}

func plus(writer http.ResponseWriter, request *http.Request) {
	execute(writer, 1)
}

func minus(writer http.ResponseWriter, request *http.Request) {
	execute(writer, -1)
}

func execute(writer http.ResponseWriter, value int) {
	number += value
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json := "{\"Value\": " + strconv.Itoa(number) + "}"
	fmt.Fprintf(writer, json)
}
