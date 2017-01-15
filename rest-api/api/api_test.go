package api_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
		log.Fatal("REST API tests unable to connect to http://rest-api:3000/api after 5 seconds!")
	}
}

// Will GET forever until it succeeds, then write true to the passed channel
func getUntilSuccess(ch chan<- bool) {
	res, err := http.Get("http://rest-api:3000/api")
	if err != nil || res == nil || res.StatusCode != 200 {
		// If we don't get 200 OK, wait a bit and try again
		time.Sleep(time.Millisecond * 500)
		getUntilSuccess(ch)
	}
	ch <- true
}

func TestRoot(t *testing.T) {
	res, err := http.Get("http://rest-api:3000/api")
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
	body := string(bodyBytes)
	expected := `{
  "_links": {
    "self": {
      "href": "/api"
    },
	"books": {
      "href": "/api/books"
	}
  }
}`
	assert.JSONEq(t, expected, body)
}

func Test404(t *testing.T) {
	request, err := http.NewRequest("GET", "http://rest-api:3000/api/url-does-not-exist", nil)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 404 {
		t.Errorf("404 expected: %d", res.StatusCode)
	}
}

// TODO: remove all books stuff and switch to coffee/orders for restbucks
func TestBooks(t *testing.T) {
	// TODO: stage some book data and assert that at least this data is present (ignore additional data?)
	// if this is a paginated API this may be difficult
	res, err := http.Get("http://rest-api:3000/api/books")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("200 expected: %d", res.StatusCode)
	}
	defer res.Body.Close()
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	body := string(bodyBytes)
	expected := `{
  "_embedded": {
    "books": [
      {
        "_links": {
          "self": {
            "href": "/api/books/ASDFAS23234"
          }
        },
        "author": "George R.R. Martin",
        "isbn": "ASDFAS23234",
        "price": 32.3,
        "title": "A Dance With Dragons"
      },
      {
        "_links": {
          "self": {
            "href": "/api/books/HJKL9898"
          }
        },
        "author": "Stieg Larsson",
        "isbn": "HJKL9898",
        "price": 9.99,
        "title": "The Girl With the Dragon Tattoo"
      }
    ]
  },
  "_links": {
    "self": {
      "href": "/api/books"
    }
  }
}`
	assert.JSONEq(t, expected, body)
}
