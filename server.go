package movingwindow

import (
	"fmt"
	"net/http"
)

type RequestsStore interface {
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
	fmt.Fprint(w, "20")
}
