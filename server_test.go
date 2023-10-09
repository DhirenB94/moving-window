package movingwindow_test

import (
	"encoding/json"
	"io"
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

		got := getReqsInLastMinFromResponse(t, response.Body)
		assertEqual(t, got.RequestsInLastMin, 1)
		assertStatus(t, response.Code, http.StatusOK)
		assertContentType(t, response, movingwindow.JSONContentType)
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

		got := getReqsInLastMinFromResponse(t, response.Body)
		assertEqual(t, got.RequestsInLastMin, 0)
		assertStatus(t, response.Code, http.StatusOK)
		assertContentType(t, response, movingwindow.JSONContentType)
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
		got := getReqsInLastMinFromResponse(t, response.Body)
		assertEqual(t, got.RequestsInLastMin, 0)
		assertStatus(t, response.Code, http.StatusOK)
		assertContentType(t, response, movingwindow.JSONContentType)

		server.ServeHTTP(response, request)
		got2 := getReqsInLastMinFromResponse(t, response.Body)
		assertEqual(t, got2.RequestsInLastMin, 1)
		assertStatus(t, response.Code, http.StatusOK)
		assertContentType(t, response, movingwindow.JSONContentType)

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

		got := getReqsInLastMinFromResponse(t, response.Body)
		assertEqual(t, got.RequestsInLastMin, 2)
		assertStatus(t, response.Code, http.StatusOK)
		assertContentType(t, response, movingwindow.JSONContentType)
	})
}

func assertEqual(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func getReqsInLastMinFromResponse(t testing.TB, body io.Reader) (reqsInLastMin movingwindow.ReqsInLastMin) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&reqsInLastMin)
	if err != nil {
		t.Fatalf("Unable to parse response from server %q into ReqsInLastMin, '%v'", body, err)
	}
	return reqsInLastMin
}
func assertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.Result().Header)
	}
}
