package request

import "net/http"

type Client interface {
	New(url string, w http.ResponseWriter, r *http.Request)
}

type ClientImpl struct {
	HTTPClient *http.Client
}
