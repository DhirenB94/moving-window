package movingwindow

import (
	"fmt"
	"net/http"
)

func RequestServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}
