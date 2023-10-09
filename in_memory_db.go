package movingwindow

import "time"

type InMemDB struct {
}

func NewInMemDB() *InMemDB {
	return &InMemDB{}
}

func (im *InMemDB) GetReqsInLastMin(reqTime time.Time) int {
	return 0
}

func (im *InMemDB) AddReqToCount(reqTime time.Time) {
}

func (im *InMemDB) GetCurrentTime() time.Time {
	return time.Now()
}
