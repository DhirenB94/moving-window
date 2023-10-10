package movingwindow

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type FileSystem struct {
	dataSource *os.File
	data       AllData
}

func NewFileSystem(dataSource *os.File) (*FileSystem, error) {
	dataSource.Seek(0, 0)
	data, err := NewData(dataSource)
	if err != nil {
		return nil, fmt.Errorf("unable to load data from the file %s, %v", dataSource.Name(), err)
	}
	return &FileSystem{
		dataSource: dataSource,
		data:       data,
	}, nil
}

func (f *FileSystem) GetReqsInLastMin(reqSecond int) int {
	data := f.data

	requestsinlastmin := 0
	oneMinAgo := reqSecond - 60

	for _, v := range data {
		if v.Second >= oneMinAgo && v.Second < reqSecond {
			requestsinlastmin += v.Count
		}
	}
	return requestsinlastmin
}
func (f *FileSystem) AddReqToCount(reqSecond int) {
	foundReq := f.data.Find(reqSecond)

	if foundReq != nil {
		foundReq.Count++
	} else {
		f.data = append(f.data, Data{
			Second: reqSecond,
			Count:  1,
		})
	}

	f.dataSource.Seek(0, 0)
	json.NewEncoder(f.dataSource).Encode(f.data)
}

func (f *FileSystem) GetCurrentSecond() int {
	currentSecond := int(time.Now().Unix())
	return currentSecond
}
func (f *FileSystem) GetAllReqs() AllData {
	return f.data
}
