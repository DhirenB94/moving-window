# Moving Window

Using only the standard library, create a Go HTTP server that on each request responds with a counter of the total number of requests that it has received during the previous 60 seconds (moving window). The server should continue to return the correct numbers after restarting it, by persisting data to a file.

## Kick Off
[Kick-Off](https://github.com/DhirenB94/moving-window/issues/1) with a task breakdown can be found here

## Run Locally

Clone the project
```bash
  git clone [https://link-to-project](https://github.com/DhirenB94/moving-window.git)
```
Go to the project directory
```bash
  cd moving-window
```
Start the server
```bash
  go run cmd/web/main.go
```
Make requests to the server (in a new terminal)
```bash
  curl -X GET http://localhost:5001
```

## Need to Know
Everytime a request is made, the server will respond with the number of requests made in the last minute.

This does NOT include the current request. 

This means the first request to the server will return a count of zero, but the second request (as long as it is made within a minute of the first) will return a count of one.

## Next Steps
1. Improve integration test - right now this tests that the count is increased everytime a request is made. Putting in a minutes sleep in order to test that only request counts from the last minute are returned, is not practical. Perhaps I could introduce some sort of time window, that can be changed to a much shorter time for testing purposes.

2. Cater for a Delete Scenario - currently there is no need to delete any of the information on the file, only add to it. 
