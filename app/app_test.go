package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApp(t *testing.T) {
	server := httptest.NewServer(App())
	defer server.Close()
	res, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code: %d", res.StatusCode)
	}
}
