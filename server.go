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
	http.Handler
}

func NewRequestServer(reqStore RequestsStore) *RequestsServer {
	rs := new(RequestsServer)
	rs.reqStore = reqStore

	router := http.NewServeMux()
	router.Handle("/", http.HandlerFunc(rs.reqCountsHandler))
	rs.Handler = router

	return rs

}

func (rs *RequestsServer) reqCountsHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := rs.reqStore.GetCurrentTime()
	reqsInLastMin := rs.reqStore.GetReqsInLastMin(currentTime)
	fmt.Fprint(w, reqsInLastMin)
	rs.reqStore.AddReqToCount(currentTime)
}
