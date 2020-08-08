package main

import "github.com/3ddiamantes/asimov-backend/internal/http"

func main() {
	router := http.InitRouter()
	router.Run(":8080")
}
