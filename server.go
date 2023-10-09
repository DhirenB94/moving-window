package movingwindow

import (
	"fmt"
	"net/http"
)

type RequestsStore interface {
	GetReqsInLastMin(reqSecond int) int
	AddReqToCount(reqSecond int)
	GetCurrentSecond() int
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
	currentSecond := rs.reqStore.GetCurrentSecond()
	reqsInLastMin := rs.reqStore.GetReqsInLastMin(currentSecond)
	fmt.Fprint(w, reqsInLastMin)
	rs.reqStore.AddReqToCount(currentSecond)
}
