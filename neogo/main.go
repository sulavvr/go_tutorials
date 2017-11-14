package main

import (
	// "fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(writer http.ResponseWriter, reader *http.Request) {
	writer.Write([]byte("Hello, World"))
}
