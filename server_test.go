package movingwindow_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	movingwindow "dhiren.brahmbhatt/moving-window"
)

const TestCurrentSecond = 1000

type StubReqStore struct {
	reqSecondCount map[int]int
}

func (s *StubReqStore) GetReqsInLastMin(reqSecond int) int {
	requestsInLastMin := 0

	//get the begining of the last minute
	oneMinAgo := reqSecond - 60

	// Iterate through the requestCount map and if the second is after oneMinAgo, add those values for those seconds
	for second, count := range s.reqSecondCount {
		if second >= oneMinAgo {
			requestsInLastMin += count
		}
	}
	return requestsInLastMin
}

func (s *StubReqStore) AddReqToCount(reqSecond int) {
	if s.reqSecondCount[reqSecond] == 0 {
		s.reqSecondCount[reqSecond] = 1
	} else {
		s.reqSecondCount[reqSecond]++
	}
}

func (s *StubReqStore) GetCurrentSecond() int {
	return TestCurrentSecond
}

func TestServer(t *testing.T) {
	t.Run("gets a count of 1, given 1 request in the last minute", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		thirtySecondsAgo := TestCurrentSecond - 30

		reqStore := &StubReqStore{
			reqSecondCount: map[int]int{
				thirtySecondsAgo: 1,
			},
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

		reqStore := &StubReqStore{
			reqSecondCount: map[int]int{
				0: 0,
			},
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

		reqStore := &StubReqStore{
			reqSecondCount: map[int]int{
				TestCurrentSecond: 0,
			},
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

		thirtySecondsAgo := TestCurrentSecond - 30
		sixtySecondsAgo := TestCurrentSecond - 60
		ninetySecondsAgo := TestCurrentSecond - 90

		reqStore := &StubReqStore{
			reqSecondCount: map[int]int{
				TestCurrentSecond: 0,
				thirtySecondsAgo:  1,
				sixtySecondsAgo:   1,
				ninetySecondsAgo:  1,
			},
		}

		server := movingwindow.NewRequestServer(reqStore)

		server.ServeHTTP(response, request)

		want := "2"
		got := response.Body.String()

		assertBody(t, got, want)
	})
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
