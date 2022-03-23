package main

import (
	"fmt"

	"first/pkg/api/http"
)

func main() {
	fmt.Println("Starting the application...")
	http.Server()
}
