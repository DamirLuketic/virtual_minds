package request

import (
	"fmt"
	"net/http"
)

func (c *ClientImpl) New(url string, w http.ResponseWriter, r *http.Request) {
	req, err := c.newRequest(http.MethodPost, url, r.Body)
	if err != nil {
		fmt.Errorf("ClientImpl.New. Error on creating a request. Error: %s", err.Error())
	}
	err = c.do(req)
	if err != nil {
		fmt.Errorf("ClientImpl.New. Error on executing a request. Error: %s", err.Error())
	}
	// Logs used for internal error handling (mock)
	w.WriteHeader(http.StatusOK)
}
