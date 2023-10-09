package movingwindow_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	movingwindow "dhiren.brahmbhatt/moving-window"
)

func TestRetrieveRequestCountAndAddCurrentRequest(t *testing.T) {
	store := movingwindow.NewInMemDB()

	server := movingwindow.NewRequestServer(store)

	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)
	assertBody(t, response.Body.String(), "0")

	server.ServeHTTP(response, request)
	assertBody(t, response.Body.String(), "01")

	server.ServeHTTP(response, request)
	assertBody(t, response.Body.String(), "012")

	assertStatus(t, response.Code, http.StatusOK)
}
