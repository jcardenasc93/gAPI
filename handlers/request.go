package handlers

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	HTTPGet    string = http.MethodGet
	HTTPPost   string = http.MethodPost
	HTTPPut    string = http.MethodPut
	HTTPDelete string = http.MethodDelete
)
const defaultVerb = HTTPGet

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
	ioReader    io.Reader
	contentType string
}

func new(url string, contentType string) *RequestHandler {
	return &RequestHandler{
		url:         url,
		contentType: contentType,
	}
}

func newReq(method string, h *RequestHandler) (r *http.Request) {
	r, err := http.NewRequest(method, h.url, h.ioReader)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return
}

func MakeRequest(url, contentType, method string) (resp *http.Response, err error) {
	handler := new(url, contentType)
	client := &http.Client{}
	m, ok := strings[method]
	if ok != true {
		return resp, errors.New("Invalid HTTP verb")
	}
	req := newReq(m, handler)
	req.Header.Set("Content-type", handler.contentType)
	resp, err = client.Do(req)

	if err != nil {
		fmt.Printf("Something get wrong. %s\n", err.Error())
		return resp, err
	}
	return
}
