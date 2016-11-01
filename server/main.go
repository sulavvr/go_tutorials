package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct {
}

func main() {
	http.Handle("/", new(MyHandler))
	http.ListenAndServe(":3000", nil)
}

func (handler *MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content Type", "text/html")
	fmt.Println(request.URL.Path[1:])

	http.ServeFile(writer, request, request.URL.Path[1:])

}
