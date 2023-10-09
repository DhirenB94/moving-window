package movingwindow_test

import (
	"bytes"
	"reflect"
	"testing"

	movingwindow "dhiren.brahmbhatt/moving-window"
)

func TestFileServer(t *testing.T) {
	t.Run("read data form a reader correctly", func(t *testing.T) {
		data := []byte(`[
			{"second":1000, "count":1},
			{"second":1030, "count":1}]`)
		var buf bytes.Buffer
		buf.Write(data)

		store := movingwindow.NewFileSystem(&buf)

		got, err := store.GetAllReqs()
		assertNoError(t, err)

		want := []movingwindow.Data{
			{Second: 1000, Count: 1},
			{Second: 1030, Count: 1},
		}
		assertData(t, got, want)
	})
	t.Run("add data to a writer correctly", func(t *testing.T) {
		data := []byte(`[
			{"second":930, "count":1},
			{"second":960, "count":1}]`)
		var buf bytes.Buffer
		//add initial data
		buf.Write(data)

		store := movingwindow.NewFileSystem(&buf)

		//add new data
		store.AddReqToCount(TestCurrentSecond)

		//check new data has been written successfully
		got, err := store.GetAllReqs()
		assertNoError(t, err)

		want := []movingwindow.Data{
			{Second: 930, Count: 1},
			{Second: 960, Count: 1},
			{Second: 1000, Count: 1},
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
func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}
