package main

import "github.com/luxarts/asimov-backend/internal/http"

func main() {
	router := http.InitRouter()
	router.Run(":8080")
}
