package handlers

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

type httpVerb int

const (
	HTTPVerbGet httpVerb = iota
	HTTPVerbPost
	HTTPVerbPut
	HTTPVerbDelete
)
const defaultVerb = "GET"

var httpVerbs map[string]httpVerb = map[string]httpVerb{
	"GET":    HTTPVerbGet,
	"POST":   HTTPVerbPost,
	"PUT":    HTTPVerbPut,
	"DELETE": HTTPVerbDelete,
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

func MakeRequest(url, contentType, verb string) (resp *http.Response, err error) {
	handler := new(url, contentType)
	client := &http.Client{}
	v, ok := httpVerbs[verb]
	if ok != true {
		return resp, errors.New("Invalid HTTP verb")
	}
	switch v {
	case HTTPVerbGet:
		resp, err = client.Get(handler.url)
	case HTTPVerbPost:
		resp, err = client.Post(handler.url, handler.contentType, handler.ioReader)
	default:
		return resp, errors.New("Invalid HTTP verb")
	}
	if err != nil {
		fmt.Printf("Something get wrong. %s\n", err.Error())
		return resp, err
	}
	return
}
