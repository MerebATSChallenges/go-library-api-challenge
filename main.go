package main

import (
	"go-library-api/api"
	"fmt"
	"net/http"
)

func main() {
	r := api.SetupRouter()

	port := ":8080"
	fmt.Println("Server running on http://localhost" + port)
	http.ListenAndServe(port, r)
}