package ui

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/jcardenasc93/gapi/handlers"
)

const httpVersion string = "1.1"

type PrettyPrinter struct {
	req  *http.Request
	resp *http.Response
}

func NewPPrinter(h *handlers.RequestHandler) *PrettyPrinter {
	return &PrettyPrinter{
		req:  h.Req,
		resp: h.Resp,
	}
}

func (pp *PrettyPrinter) Print() {
	pp.printReq()
	pp.printResp()
}

func (pp *PrettyPrinter) printReq() {
	fmt.Printf("\r\n%s %s HTTP/%s\r\n", pp.req.Method, pp.req.URL.RequestURI(), httpVersion)
	fmt.Printf("%s %s\r\n", "Host:", pp.req.Host)
	pp.prettyHeaders(pp.req.Header)
	body := pp.req.Body
	if body != nil {
		pp.prettyJson(body)
		defer body.Close()
	}
}

func (pp *PrettyPrinter) printResp() {
	pp.prettyHeaders(pp.resp.Header)
	body := pp.resp.Body
	if body != nil {
		pp.prettyJson(body)
		defer body.Close()
	}
}

func (pp *PrettyPrinter) prettyHeaders(headers map[string][]string) {
	for k, v := range headers {
		vv := strings.Join(v, ", ")
		fmt.Printf("%s: %s\n", k, vv)
	}
	fmt.Println()
}

func (pp *PrettyPrinter) prettyJson(rawBody io.ReadCloser) {
	var (
		body []byte
		buff bytes.Buffer
	)
	body, err := ioutil.ReadAll(rawBody)
	if err != nil {
		fmt.Printf("Something get wrong. %s\n", err.Error())
		os.Exit(1)
	}
	json.Indent(&buff, body, "", "\t")
	fmt.Println(string(buff.Bytes()))
}
