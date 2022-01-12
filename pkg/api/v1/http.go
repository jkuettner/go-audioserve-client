package apiv1

import (
	"net/http"
	"time"
)

func newHTTPHandler(timeout time.Duration) *httpHandler {
	return &httpHandler{
		c: http.Client{
			Timeout: timeout,
		},
	}
}

type httpHandler struct {
	c http.Client
}

func (h *httpHandler) Do(req *http.Request) (*http.Response, error) {
	req.Header["Content-Type"] = []string{"application/json"}
	req.Header["accept"] = []string{"application/json"}
	return h.c.Do(req)
}
