package movingwindow

import (
	"encoding/json"
	"fmt"
	"io"
)

type ReqsInLastMin struct {
	RequestsInLastMin int `json:"requestsinlastmin"`
}
type Data struct {
	Second int `json:"second"`
	Count  int `json:"count"`
}

func NewData(reader io.Reader) ([]Data, error) {
	var data []Data
	err := json.NewDecoder(reader).Decode(&data)
	if err != nil {
		err = fmt.Errorf("problem parsing data, %v", err)
	}
	return data, err
}
