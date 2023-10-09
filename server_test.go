package movingwindow_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	movingwindow "dhiren.brahmbhatt/moving-window"
)

type StubReqStore struct {
	count int
}

func (s *StubReqStore) GetReqsInLastMin(reqTime time.Time) int {
	return s.count
}
func (s *StubReqStore) AddReqToCount(reqTime time.Time) {
	s.count++
}

func Test(t *testing.T) {
	t.Run("gets the count", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		reponse := httptest.NewRecorder()

		reqStore := &StubReqStore{
			count: 20,
		}
		server := movingwindow.NewRequestServer(reqStore)

		server.ServeHTTP(reponse, request)

		want := "20"
		got := reponse.Body.String()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("adds one to the count for the current request", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		reponse := httptest.NewRecorder()

		reqStore := &StubReqStore{
			count: 20,
		}
		server := movingwindow.NewRequestServer(reqStore)

		server.ServeHTTP(reponse, request)
		server.ServeHTTP(reponse, request)

		want := "2021"
		got := reponse.Body.String()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
