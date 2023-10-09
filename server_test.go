package movingwindow_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	movingwindow "dhiren.brahmbhatt/moving-window"
)

const timeNow = "2023-10-09 14:30:00"
const thirtySecsAgo = "2023-10-09 14:29:30"
const fourtySecsAgo = "2023-10-09 14:29:20"
const ninetySecsAgo = "2023-10-09 14:28:30"

type StubReqStore struct {
	reqTimeCount map[time.Time]int
	time         string
}

func (s *StubReqStore) GetReqsInLastMin(reqTime time.Time) int {
	requestsInLastMin := 0

	//get the begining of the last minute
	lastMinute := reqTime.Add(-61 * time.Second)

	// Iterate through the requestCount map and if the time is after the last minute starts, sum those values
	for t, count := range s.reqTimeCount {
		if t.After(lastMinute) {
			requestsInLastMin += count
		}
	}
	return requestsInLastMin
}

func (s *StubReqStore) AddReqToCount(reqTime time.Time) {
	if s.reqTimeCount[reqTime] == 0 {
		s.reqTimeCount[reqTime] = 1
	} else {
		s.reqTimeCount[reqTime]++
	}
}

func (s *StubReqStore) GetCurrentTime() time.Time {
	return timeParser(s.time)
}

func Test(t *testing.T) {
	t.Run("gets a count of 1, given 1 request in the last minute", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		parsedTimeThirtySecsAgo := timeParser(thirtySecsAgo)

		reqStore := &StubReqStore{
			reqTimeCount: map[time.Time]int{
				parsedTimeThirtySecsAgo: 1,
			},
			time: timeNow,
		}
		server := movingwindow.NewRequestServer(reqStore)

		server.ServeHTTP(response, request)

		want := "1"
		got := response.Body.String()

		assertBody(t, got, want)
	})
	t.Run("gets a count of 0, given no requests in the last minute", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		parsedTime := timeParser(timeNow)

		reqStore := &StubReqStore{
			reqTimeCount: map[time.Time]int{
				parsedTime: 0,
			},
			time: timeNow,
		}
		server := movingwindow.NewRequestServer(reqStore)

		server.ServeHTTP(response, request)

		want := "0"
		got := response.Body.String()

		assertBody(t, got, want)
	})
	t.Run("adds one to the count for the current request", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		parsedTime := timeParser(timeNow)

		reqStore := &StubReqStore{
			reqTimeCount: map[time.Time]int{
				parsedTime: 0,
			},
			time: timeNow,
		}
		server := movingwindow.NewRequestServer(reqStore)

		server.ServeHTTP(response, request)
		server.ServeHTTP(response, request)

		want := "01" //0, 1
		got := response.Body.String()

		assertBody(t, got, want)
	})
	t.Run("only get the count for requests within the last minute", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		parseTimeNow := timeParser(timeNow)
		parsedthirtySecsAgo := timeParser(thirtySecsAgo)
		parsedfourtySecsAgo := timeParser(fourtySecsAgo)
		parsedninetySecsAgo := timeParser(ninetySecsAgo)

		reqStore := &StubReqStore{
			reqTimeCount: map[time.Time]int{
				parseTimeNow:        0,
				parsedthirtySecsAgo: 1,
				parsedfourtySecsAgo: 1,
				parsedninetySecsAgo: 1,
			},
			time: timeNow,
		}

		server := movingwindow.NewRequestServer(reqStore)

		server.ServeHTTP(response, request)

		want := "2"
		got := response.Body.String()

		assertBody(t, got, want)
	})
}

func timeParser(input string) time.Time {
	layout := "2006-01-02 15:04:05"
	parsedTime, _ := time.Parse(layout, input)
	return parsedTime
}

func assertBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
