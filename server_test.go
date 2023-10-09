package movingwindow_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	movingwindow "dhiren.brahmbhatt/moving-window"
)

const testTime = "2023-10-09 14:30:00"

type StubReqStore struct {
	reqTimeCount map[time.Time]int
}

func (s *StubReqStore) GetReqsInLastMin(reqTime time.Time) int {
	return s.reqTimeCount[reqTime]
}
func (s *StubReqStore) AddReqToCount(reqTime time.Time) {
	s.reqTimeCount[reqTime]++
}

func (s *StubReqStore) GetCurrentTime() time.Time {
	return timeParser(testTime)
}

func Test(t *testing.T) {
	t.Run("gets the count", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		reponse := httptest.NewRecorder()

		parsedTime := timeParser(testTime)

		reqStore := &StubReqStore{
			reqTimeCount: map[time.Time]int{
				parsedTime: 1,
			},
		}
		server := movingwindow.NewRequestServer(reqStore)

		server.ServeHTTP(reponse, request)

		want := "1"
		got := reponse.Body.String()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("adds one to the count for the current request", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		reponse := httptest.NewRecorder()

		parsedTime := timeParser(testTime)
		reqStore := &StubReqStore{
			reqTimeCount: map[time.Time]int{
				parsedTime: 10,
			},
		}
		server := movingwindow.NewRequestServer(reqStore)

		server.ServeHTTP(reponse, request)
		server.ServeHTTP(reponse, request)

		want := "1011"
		got := reponse.Body.String()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func timeParser(input string) time.Time {
	layout := "2006-01-02 15:04:05"
	parsedTime, _ := time.Parse(layout, input)
	return parsedTime
}
