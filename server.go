package movingwindow

import (
	"fmt"
	"net/http"
	"time"
)

type RequestsStore interface {
	GetReqsInLastMin(reqTime time.Time) int
	AddReqToCount(reqTime time.Time)
	GetCurrentTime() time.Time
}

type RequestsServer struct {
	reqStore RequestsStore
}

func NewRequestServer(reqStore RequestsStore) *RequestsServer {
	return &RequestsServer{
		reqStore: reqStore,
	}
}

func (rs *RequestsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	currentTime := rs.reqStore.GetCurrentTime()
	reqsInLastMin := rs.reqStore.GetReqsInLastMin(currentTime)
	fmt.Fprint(w, reqsInLastMin)
	rs.reqStore.AddReqToCount(currentTime)
}
