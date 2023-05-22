package main

import (
	"fmt"
	"net/http"

	"github.com/jeroendee/go-demo-web-api/internal/webapi"
)

func main() {
	port := "3000"
	s := webapi.NewServer()
	fmt.Printf("Server starting on localhost:%v\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), s)
}
