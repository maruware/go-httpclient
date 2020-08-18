package httpclient_test

import (
	"bytes"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/maruware/go-httpclient"
)

type ResponseData struct {
	URL string `json:"url"`
}

func TestGet(t *testing.T) {
	base, err := url.Parse("https://postman-echo.com")
	if err != nil {
		t.Fatalf("Failed to parse url")
	}
	c := httpclient.HttpClient{
		BaseURL: base,
	}
	c.Timeout = time.Second * 10

	res, err := c.Get("/get", nil)
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
	opt := httpclient.RequestOption{
		Query: "key=example",
	}
	res, err := c.Post("/post", httpclient.ContentTypeJson, body, &opt)
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

func TestGetJson(t *testing.T) {
	base, err := url.Parse("https://postman-echo.com")
	if err != nil {
		t.Fatalf("Failed to parse url")
	}
	c := httpclient.HttpClient{
		BaseURL: base,
	}
	c.Timeout = time.Second * 10

	var d ResponseData
	err = c.GetJson("/get", &d, nil)
	if err != nil {
		t.Fatalf("Failed to get: %v", err)
	}
	if len(d.URL) <= 0 {
		t.Fatalf("Bad response data: %+v", d)
	}
}

func TestPostJson(t *testing.T) {
	base, err := url.Parse("https://postman-echo.com")
	if err != nil {
		t.Fatalf("Failed to parse url")
	}

	c := httpclient.HttpClient{
		BaseURL: base,
	}
	c.Timeout = time.Second * 10

	reqd := map[string]string{"name": "taro"}
	var resd ResponseData
	err = c.PostJson("/post", reqd, &resd, nil)
	if err != nil {
		t.Fatalf("Failed to post: %v", err)
	}

	if len(resd.URL) <= 0 {
		t.Fatalf("Bad response data: %+v", resd)
	}
}
