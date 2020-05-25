package httpclient_test

import (
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/maruware/go-httpclient"
)

func TestGet(t *testing.T) {
	base, err := url.Parse("https://google.com")
	if err != nil {
		t.Fatalf("Failed to parse url")
	}
	c := httpclient.HttpClient{
		BaseURL: base,
	}
	c.Timeout = time.Second * 10

	res, err := c.Get("/doodles/")
	if err != nil {
		t.Fatalf("Failed to get doodles: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("expect status %d but %d", http.StatusOK, res.StatusCode)
	}
}
