package test

import (
	"asimov-backend/internal/defines"
	router "asimov-backend/internal/http"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var prevJwtSecret string

// Tests
func TestMain(m *testing.M) {
	setup()
	fmt.Println("Running tests")
	code := m.Run()
	fmt.Println("Tests run")
	teardown()
	os.Exit(code)
}

// Utils
func setup() {
	router.InitRouter()
	prevJwtSecret = os.Getenv(defines.EnvJWTSecret)
	os.Setenv(defines.EnvJWTSecret, "dGVzdA")
}
func teardown() {
	os.Setenv(defines.EnvJWTSecret, prevJwtSecret)
	prevJwtSecret = ""
}
func executeRequest(method string, url string, headers map[string]string, body string) *httptest.ResponseRecorder {
	var request *http.Request

	if body != "" {
		request, _ = http.NewRequest(method, url, strings.NewReader(body))

	} else {
		request, _ = http.NewRequest(method, url, nil)
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	return response
}
