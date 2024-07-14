package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:    ":3000",
		Handler: http.HandlerFunc(basicHandler),
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to listen with error: ", err)
	}
}

func basicHandler(w http.ResponseWriter, t *http.Request) {
	w.Write([]byte("Hello there!"))
}
