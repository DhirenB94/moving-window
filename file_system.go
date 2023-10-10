package movingwindow

import (
	"encoding/json"
	"os"
)

type FileSystem struct {
	dataSource *os.File
}

func NewFileSystem(dataSource *os.File) *FileSystem {
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

	f.dataSource.Seek(0, 0)
	json.NewEncoder(f.dataSource).Encode(data)
}
func (f *FileSystem) GetCurrentSecond() int {
	return 0
}
func (f *FileSystem) GetAllReqs() []Data {
	f.dataSource.Seek(0, 0)
	data, _ := NewData(f.dataSource)
	return data
}
