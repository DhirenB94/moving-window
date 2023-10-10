package movingwindow

import (
	"encoding/json"
	"io"
)

type FileSystem struct {
	dataSource io.ReadWriter
}

func NewFileSystem(dataSource io.ReadWriter) *FileSystem {
	return &FileSystem{dataSource: dataSource}
}

func (f *FileSystem) GetReqsInLastMin(reqSecond int) int {
	return 0
}
func (f *FileSystem) AddReqToCount(reqSecond int) {
	data := f.GetAllReqs()

	data = append(data, Data{
		Second: reqSecond,
		Count:  1,
	})

	json.NewEncoder(f.dataSource).Encode(data)
}
func (f *FileSystem) GetCurrentSecond() int {
	return 0
}
func (f *FileSystem) GetAllReqs() []Data {
	data, _ := NewData(f.dataSource)
	return data
}
