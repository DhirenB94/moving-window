package movingwindow

import (
	"sync"
	"time"
)

type InMemDB struct {
	reqSecondCount map[int]int
	mutex          sync.Mutex
}

func NewInMemDB() *InMemDB {
	return &InMemDB{
		reqSecondCount: map[int]int{
			0: 0,
		},
		mutex: sync.Mutex{},
	}
}

func (im *InMemDB) GetReqsInLastMin(reqSecond int) int {
	im.mutex.Lock()
	defer im.mutex.Unlock()

	requestsInLastMin := 0
	// get the begining of the last minute
	oneMinAgo := reqSecond - 60

	// Iterate through the requestCount map and if the time is after the last minute starts, sum those values
	for second, count := range im.reqSecondCount {
		if second >= oneMinAgo {
			requestsInLastMin += count
		}
	}

	return requestsInLastMin
}

func (im *InMemDB) AddReqToCount(reqSecond int) {
	//Update the requestCount map for the current second
	im.mutex.Lock()
	defer im.mutex.Unlock()
	im.reqSecondCount[reqSecond]++
}

func (im *InMemDB) GetCurrentSecond() int {
	return int(time.Now().Unix())
}
