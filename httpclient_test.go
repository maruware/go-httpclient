package httpclient_test

import (
	"bytes"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/maruware/go-httpclient"
)

func TestGet(t *testing.T) {
	base, err := url.Parse("https://postman-echo.com")
	if err != nil {
		t.Fatalf("Failed to parse url")
	}
	c := httpclient.HttpClient{
		BaseURL: base,
	}
	c.Timeout = time.Second * 10

	res, err := c.Get("/get")
	if err != nil {
		t.Fatalf("Failed to get: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("expect status %d but %d", http.StatusOK, res.StatusCode)
	}
}

func TestPost(t *testing.T) {
	base, err := url.Parse("https://postman-echo.com")
	if err != nil {
		t.Fatalf("Failed to parse url")
	}

	c := httpclient.HttpClient{
		BaseURL: base,
	}
	c.Timeout = time.Second * 10

	data := map[string]string{"name": "taro"}
	body := bytes.NewBuffer(nil)
	httpclient.EncodeJson(data, body)
	res, err := c.Post("/post", httpclient.ContentTypeJson, body)
	if err != nil {
		t.Fatalf("Failed to post: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("expect status %d but %d", http.StatusOK, res.StatusCode)
	}

	var d interface{}
	err = httpclient.DecodeJson(&d, res.Body)
	defer res.Body.Close()
	if err != nil {
		t.Fatalf("Failed to decode body")
	}
}
