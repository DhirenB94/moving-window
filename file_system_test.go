package movingwindow_test

import (
	"reflect"
	"strings"
	"testing"

	movingwindow "dhiren.brahmbhatt/moving-window"
)

func TestFileServer(t *testing.T) {
	t.Run("read data form a reader correctly", func(t *testing.T) {
		data := strings.NewReader(`[
			{"second":1000, "count":1},
			{"second":1030, "count":1}]`)

		store := movingwindow.NewFileSystem(data)

		got := store.GetAllReqs()
		want := []movingwindow.Data{
			{Second: 1000, Count: 1},
			{Second: 1030, Count: 1},
		}
		assertData(t, got, want)
	})
}

func assertData(t testing.TB, got, want []movingwindow.Data) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
