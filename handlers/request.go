package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

const (
	HTTPGet    string = http.MethodGet
	HTTPPost   string = http.MethodPost
	HTTPPut    string = http.MethodPut
	HTTPDelete string = http.MethodDelete
)
const DefaultVerb = HTTPGet

var strings map[string]string = map[string]string{
	"GET":    HTTPGet,
	"get":    HTTPGet,
	"POST":   HTTPPost,
	"post":   HTTPPost,
	"PUT":    HTTPPut,
	"put":    HTTPPut,
	"DELETE": HTTPDelete,
	"delete": HTTPDelete,
}

type RequestHandler struct {
	url         string
	contentType string
	Req         *http.Request
	Resp        *http.Response
}

func new(url string, contentType string) *RequestHandler {
	return &RequestHandler{
		url:         url,
		contentType: contentType,
	}
}

func newReq(method string, h *RequestHandler) (r *http.Request) {
	r, err := http.NewRequest(method, h.url, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return
}

func MakeRequest(url, contentType, method string) (handler *RequestHandler, err error) {
	handler = new(url, contentType)
	client := &http.Client{}
	m, ok := strings[method]
	if ok != true {
		return nil, errors.New("Invalid HTTP verb")
	}
	handler.Req = newReq(m, handler)
	handler.Req.Header.Set("Content-type", handler.contentType)
	handler.Req.Header.Set("X-example", "hello from gAPI")
	handler.Req.Header.Add("X-example", "other")
	handler.Resp, err = client.Do(handler.Req)

	if err != nil {
		fmt.Printf("Something get wrong. %s\n", err.Error())
		return nil, err
	}

	return
}
