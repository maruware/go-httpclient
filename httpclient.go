package httpclient

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type HttpClient struct {
	http.Client
	BaseURL *url.URL
	Header  http.Header
}

func (c *HttpClient) newRequest(method, path string, body io.Reader) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}
	for key, values := range c.Header {
		for _, v := range values {
			req.Header.Set(key, v)
		}
	}
	return req, err
}

func (c *HttpClient) send(req *http.Request) (*http.Response, error) {
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("http status code error: %v", res.StatusCode)
	}
	return res, nil
}

func (c *HttpClient) SetHeader(key, value string) {
	c.Header.Set(key, value)
}

func (c *HttpClient) Request(method, path string, body io.Reader) (*http.Response, error) {
	req, err := c.newRequest(method, path, body)
	if err != nil {
		return nil, err
	}

	return c.send(req)
}

func (c *HttpClient) Get(path string) (*http.Response, error) {
	req, err := c.newRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	return c.send(req)
}

func (c *HttpClient) Post(path string, contentType string, body io.Reader) (*http.Response, error) {
	req, err := c.newRequest("POST", path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", contentType)
	return c.send(req)
}

func (c *HttpClient) Put(path string, contentType string, body io.Reader) (*http.Response, error) {
	req, err := c.newRequest("PUT", path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", contentType)
	return c.send(req)
}

func (c *HttpClient) Delete(path string, contentType string, body io.Reader) (*http.Response, error) {
	req, err := c.newRequest("DELETE", path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", contentType)
	return c.send(req)
}

func (c *HttpClient) GetJson(path string, v interface{}) error {
	req, err := c.newRequest("GET", path, nil)
	if err != nil {
		return err
	}
	res, err := c.send(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if err := DecodeJson(v, res.Body); err != nil {
		return err
	}
	return nil
}

func (c *HttpClient) PostJson(path string, reqData interface{}, resData interface{}) error {
	b := bytes.NewBuffer(nil)
	if err := EncodeJson(reqData, b); err != nil {
		return err
	}

	req, err := c.newRequest("POST", path, b)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", ContentTypeJson)
	res, err := c.send(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if err := DecodeJson(resData, res.Body); err != nil {
		return err
	}
	return nil
}
