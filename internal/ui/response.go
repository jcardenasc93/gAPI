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
)

const httpVersion string = "1.1"

type PrettyPrinter struct {
	req  *http.Request
	resp *http.Response
}

func NewPPrinter(req *http.Request, resp *http.Response) *PrettyPrinter {
	return &PrettyPrinter{
		req:  req,
		resp: resp,
	}
}

func (pp *PrettyPrinter) Print() {
	req := pp.printReq()
	resp := pp.printResp()

	out := fmt.Sprintf("%s\n\r%s", req, resp)
	fmt.Println(out)
}

func (pp *PrettyPrinter) printReq() string {
	var (
		uri, host, headers, body string
	)
	uri = fmt.Sprintf("\r\n%s %s HTTP/%s", pp.req.Method, pp.req.URL.RequestURI(), httpVersion)
	host = fmt.Sprintf("%s %s", "Host:", pp.req.Host)
	headers = pp.prettyHeaders(pp.req.Header)
	b := pp.req.Body
	if b != nil {
		body = prettyJson(&b)
		defer b.Close()
	}
	return fmt.Sprintf("%s\n%s\n%s\n%s", uri, host, headers, body)
}

func (pp *PrettyPrinter) printResp() string {
	var (
		headers, body string
	)
	headers = pp.prettyHeaders(pp.resp.Header)
	b := pp.resp.Body
	if b != nil {
		body = prettyJson(&b)
		defer b.Close()
	}
	return fmt.Sprintf("%s\r\n\n%s\r", body, headers)
}

func (pp *PrettyPrinter) prettyHeaders(headers map[string][]string) string {
	var output string
	for k, v := range headers {
		vv := strings.Join(v, ", ")
		header := fmt.Sprintf("%s: %s\n", k, vv)
		output = fmt.Sprintf("%s%s", output, header)
	}
	return fmt.Sprintf("%s", output)
}

func prettyJson(rawBody *io.ReadCloser) string {
	var (
		body   []byte
		buff   bytes.Buffer
		output string
	)
	body, err := ioutil.ReadAll(*rawBody)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
	json.Indent(&buff, body, "", "\t")
	output = fmt.Sprintf("%s", string(buff.Bytes()))
	return output
}
