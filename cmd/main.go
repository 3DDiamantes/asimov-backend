package main

import "asimov-backend/internal/http"

func main() {
	router := http.InitRouter()
	router.Run(":8080")
}
