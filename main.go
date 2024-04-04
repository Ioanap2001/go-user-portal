package main

import (
	"fmt"
	"net/http"
)

func main() {
	setupRoutes()
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
