package movingwindow

import (
	"sync"
	"time"
)

type InMemDB struct {
	requestCount map[time.Time]int
	mutex sync.Mutex
}

func NewInMemDB() *InMemDB {
	return &InMemDB{
		requestCount: map[time.Time]int{
			{}: 0,
		},
		mutex: sync.Mutex{},
	}
}

func (im *InMemDB) GetReqsInLastMin(reqTime time.Time) int {
		im.mutex.Lock()
		defer im.mutex.Unlock()
	
		requestsInLastMin := 0
	
		//get the begining of the last minute
		lastMinute := reqTime.Add(-61 * time.Second)
	
		// Iterate through the requestCount map and if the time is within the last minute, sum those values
		for t, count := range im.requestCount {
			if t.After(lastMinute) {
				requestsInLastMin += count
			}
		}
		return requestsInLastMin
}

func (im *InMemDB) AddReqToCount(reqTime time.Time) {
		//Update the requestCount map for the current second
		im.mutex.Lock()
		defer im.mutex.Unlock()
		im.requestCount[reqTime]++
}

func (im *InMemDB) GetCurrentTime() time.Time {
	return time.Now()
}
