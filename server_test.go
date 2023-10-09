package movingwindow_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	movingwindow "dhiren.brahmbhatt/moving-window"
)

type StubReqStore struct {
}

func Test(t *testing.T) {
	t.Run("returns dummy data for number of requests", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		reponse := httptest.NewRecorder()

		reqStore := &StubReqStore{}
		server := movingwindow.NewRequestServer(reqStore)

		server.ServeHTTP(reponse, request)

		want := "20"
		got := reponse.Body.String()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
