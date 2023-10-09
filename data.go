package movingwindow

type ReqsInLastMin struct {
	RequestsInLastMin int `json:"requestsinlastmin"`
}
type Data struct {
	Second int `json:"second"`
	Count  int `json:"count"`
}
