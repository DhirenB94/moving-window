package movingwindow

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type FileSystem struct {
	file *os.File
	data AllData
}

func NewFileSystem(file *os.File) (*FileSystem, error) {
	err := initialiseFile(file)
	if err != nil {
		return nil, fmt.Errorf("unable to initialise file %s, %v", file.Name(), err)
	}
	data, err := NewData(file)
	if err != nil {
		return nil, fmt.Errorf("unable to load data from the file %s, %v", file.Name(), err)
	}
	return &FileSystem{
		file: file,
		data: data,
	}, nil
}

func initialiseFile(file *os.File) error {
	file.Seek(0, 0)
	stats, err := file.Stat()
	if err != nil {
		return fmt.Errorf("problem getting stats from file %s, %v", file.Name(), err)
	}
	if stats.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}
	return nil
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

	f.file.Seek(0, 0)
	json.NewEncoder(f.file).Encode(f.data)
}

func (f *FileSystem) GetCurrentSecond() int {
	currentSecond := int(time.Now().Unix())
	return currentSecond
}
func (f *FileSystem) GetAllReqs() AllData {
	return f.data
}
