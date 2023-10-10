package movingwindow_test

import (
	"os"
	"reflect"
	"testing"
	"time"

	movingwindow "dhiren.brahmbhatt/moving-window"
)

func TestFileServer(t *testing.T) {
	t.Run("read data form a file correctly", func(t *testing.T) {
		tempFile := createTempFile(t, `[
			{"second":1000, "count":1},
			{"second":1030, "count":1}]`)

		defer closeTempFile(tempFile)

		store := movingwindow.NewFileSystem(tempFile)

		got := store.GetAllReqs()

		want := []movingwindow.Data{
			{Second: 1000, Count: 1},
			{Second: 1030, Count: 1},
		}
		assertData(t, got, want)
	})
	t.Run("add count for current request second, when it does not already exist", func(t *testing.T) {
		tempFile := createTempFile(t, `[
			{"second":930, "count":1},
			{"second":960, "count":1}]`)
		defer closeTempFile(tempFile)

		store := movingwindow.NewFileSystem(tempFile)

		//add new data
		store.AddReqToCount(TestCurrentSecond)

		//check new data has been written successfully
		got := store.GetAllReqs()

		want := []movingwindow.Data{
			{Second: 930, Count: 1},
			{Second: 960, Count: 1},
			{Second: 1000, Count: 1},
		}
		assertData(t, got, want)
	})
	t.Run("add count for current request second, when it does exist", func(t *testing.T) {
		tempFile := createTempFile(t, `[
			{"second":930, "count":1},
			{"second":1000, "count":1}]`)
		defer closeTempFile(tempFile)

		store := movingwindow.NewFileSystem(tempFile)

		//add new data
		store.AddReqToCount(TestCurrentSecond)

		//check new data has been written successfully
		got := store.GetAllReqs()

		want := []movingwindow.Data{
			{Second: 930, Count: 1},
			{Second: 1000, Count: 2},
		}
		assertData(t, got, want)
	})
	t.Run("get the correct second", func(t *testing.T) {
		tempFile := createTempFile(t, `[
			{"second":1000, "count":1},
			{"second":1030, "count":1}]`)

		defer closeTempFile(tempFile)

		store := movingwindow.NewFileSystem(tempFile)

		storeCurrentSecond := store.GetCurrentSecond()
		currentSecond := int(time.Now().Unix())

		assertEqual(t, storeCurrentSecond, currentSecond)
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

func createTempFile(t testing.TB, initialData string) *os.File {
	t.Helper()

	tempFile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}
	tempFile.Write([]byte(initialData))

	return tempFile
}

func closeTempFile(tempFile *os.File) {
	tempFile.Close()
	os.Remove(tempFile.Name())
}
