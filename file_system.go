package movingwindow

import (
	"encoding/json"
	"os"
	"time"
)

type FileSystem struct {
	dataSource *os.File
}

func NewFileSystem(dataSource *os.File) *FileSystem {
	return &FileSystem{dataSource: dataSource}
}

func (f *FileSystem) GetReqsInLastMin(reqSecond int) int {
	data := f.GetAllReqs()

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
	data := f.GetAllReqs()
	foundReq := data.Find(reqSecond)

	if foundReq != nil {
		foundReq.Count++
	} else {
		data = append(data, Data{
			Second: reqSecond,
			Count:  1,
		})
	}

	f.dataSource.Seek(0, 0)
	json.NewEncoder(f.dataSource).Encode(data)
}

func (f *FileSystem) GetCurrentSecond() int {
	currentSecond := int(time.Now().Unix())
	return currentSecond
}
func (f *FileSystem) GetAllReqs() AllData {
	f.dataSource.Seek(0, 0)
	data, _ := NewData(f.dataSource)
	return data
}
