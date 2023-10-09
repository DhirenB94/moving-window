package movingwindow

import (
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
func (f *FileSystem) GetAllReqs() []Data {
	var data []Data
	return data
}
