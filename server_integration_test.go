package movingwindow_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	movingwindow "dhiren.brahmbhatt/moving-window"
)

func TestRetrieveRequestCountAndAddCurrentRequest(t *testing.T) {
	tempFile := createTempFile(t, "")
	defer closeTempFile(tempFile)

	store := movingwindow.NewFileSystem(tempFile)

	server := movingwindow.NewRequestServer(store)

	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)
	got := getReqsInLastMinFromResponse(t, response.Body)
	assertEqual(t, got.RequestsInLastMin, 0)
	assertStatus(t, response.Code, http.StatusOK)

	time.Sleep(1 * time.Second)
	server.ServeHTTP(response, request)
	got2 := getReqsInLastMinFromResponse(t, response.Body)
	assertEqual(t, got2.RequestsInLastMin, 1)
	assertStatus(t, response.Code, http.StatusOK)

	time.Sleep(1 * time.Second)
	server.ServeHTTP(response, request)
	got3 := getReqsInLastMinFromResponse(t, response.Body)
	assertEqual(t, got3.RequestsInLastMin, 2)
	assertStatus(t, response.Code, http.StatusOK)

	assertStatus(t, response.Code, http.StatusOK)
}
