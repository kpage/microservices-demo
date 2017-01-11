package api_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"
)

func init() {
	waitForGin()
}

// When gin starts the app, there is sometimes a pause where it will return 502 until the underlying server
// is ready.  We wait up to 5 seconds for the server to start returning 200.
func waitForGin() {
	ch := make(chan bool, 1)
	go getUntilSuccess(ch)
	select {
	case <-ch:
		// ok!
	case <-time.After(time.Second * 5):
		log.Fatal("REST API tests unable to connect to http://rest-api:3000 after 5 seconds!")
	}
}

// Will GET forever until it succeeds, then write true to the passed channel
func getUntilSuccess(ch chan<- bool) {
	res, err := http.Get("http://rest-api:3000")
	if err != nil || res == nil || res.StatusCode != 200 {
		// If we don't get 200 OK, wait a bit and try again
		time.Sleep(time.Millisecond * 500)
		getUntilSuccess(ch)
	}
	ch <- true
}

func TestHelloWorld(t *testing.T) {
	res, err := http.Get("http://rest-api:3000")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
	defer res.Body.Close()
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	// Check the response body is what we expect.
	body := string(bodyBytes)
	expected := "Hello world!\n"
	if body != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			body, expected)
	}
}

func Test404(t *testing.T) {
	request, err := http.NewRequest("GET", "http://rest-api:3000/url-does-not-exist", nil)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 404 {
		t.Errorf("404 expected: %d", res.StatusCode)
	}
}
