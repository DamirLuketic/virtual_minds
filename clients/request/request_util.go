package request

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var httpNewRequest = http.NewRequest

func NewClient() Client {
	return &ClientImpl{
		HTTPClient: &http.Client{},
	}
}

func (c *ClientImpl) newRequest(method string, path string, body io.ReadCloser) (*http.Request, error) {
	url, err := c.parseURL(path)
	if err != nil {
		return nil, err
	}
	req, err := httpNewRequest(method, url.String(), body)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (c *ClientImpl) parseURL(path string) (*url.URL, error) {
	return url.Parse(path)
}

func (c *ClientImpl) do(req *http.Request) error {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	}
	return getError(req.URL.String(), req.Method, res.StatusCode)
}

func getError(url, method string, statusCode int) error {
	return errors.New(fmt.Sprintf(
		"Error on executing request with path: %s, and method: %s. Response StatusCode: %d.",
		url, method, statusCode))
}
