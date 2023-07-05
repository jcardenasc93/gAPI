package handler

import (
	"errors"
	"net/http"
	"strings"
	"time"
)

const (
	HTTPGet    string = http.MethodGet
	HTTPPost   string = http.MethodPost
	HTTPPut    string = http.MethodPut
	HTTPDelete string = http.MethodDelete
)
const DefaultVerb string = HTTPGet
const DefaultTimeout time.Duration = 10 * time.Second

var httpVerbs map[string]string = map[string]string{
	"GET":    HTTPGet,
	"get":    HTTPGet,
	"POST":   HTTPPost,
	"post":   HTTPPost,
	"PUT":    HTTPPut,
	"put":    HTTPPut,
	"DELETE": HTTPDelete,
	"delete": HTTPDelete,
}

type Handler struct {
	Req  *http.Request
	Resp *http.Response
}

func (h *Handler) addHeaders(headers []string) {
	for _, header := range headers {
		hVal := strings.Split(header, ":")
		k, v := hVal[0], hVal[1]
		h.Req.Header.Add(k, v)
	}
}

func MakeReq(url string, method string, headers []string) (*Handler, error) {
	c := &http.Client{Timeout: DefaultTimeout}
	method, err := checkMethod(method)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		// TODO: create error handler
		return nil, err
	}
	h := &Handler{Req: req}

	if len(headers) > 0 {
		h.addHeaders(headers)
	}

	resp, err := c.Do(req)
	if err != nil {
		// TODO: create error handler
		return nil, err
	}
	h.Resp = resp
	return h, nil
}

func checkMethod(method string) (string, error) {
	method, ok := httpVerbs[method]
	if ok == false {
		return "", errors.New("Invalid HTTP method")
	}
	return method, nil
}
