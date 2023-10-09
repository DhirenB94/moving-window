package movingwindow

import (
	"encoding/json"
	"io"
)

type FileSystem struct {
	dataSource io.Reader
}

func NewFileSystem(dataSource io.Reader) *FileSystem {
	return &FileSystem{dataSource: dataSource}
}

func (f *FileSystem) GetReqsInLastMin(reqSecond int) int {
	return 0
}
func (f *FileSystem) AddReqToCount(reqSecond int) {

}
func (f *FileSystem) GetCurrentSecond() int {
	return 0
}
func (f *FileSystem) GetAllReqs() ([]Data, error) {
	var data []Data
	err := json.NewDecoder(f.dataSource).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
