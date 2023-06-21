package ui

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type PrettyPrinter struct {
	resp *http.Response
}

func NewPPrinter(resp *http.Response) *PrettyPrinter {
	return &PrettyPrinter{
		resp: resp,
	}
}

func (pp *PrettyPrinter) PrintResp() {
	pp.prettyHeaders()
	pp.prettyJson()
}

func (pp *PrettyPrinter) prettyHeaders() {
	for k, v := range pp.resp.Header {
		vv := strings.Join(v, ", ")
		fmt.Printf("%s: %s\n", k, vv)
	}
	fmt.Println()
}

func (pp *PrettyPrinter) prettyJson() {
	var (
		body []byte
		buff bytes.Buffer
	)
	defer pp.resp.Body.Close()
	body, err := ioutil.ReadAll(pp.resp.Body)
	if err != nil {
		fmt.Printf("Something get wrong. %s\n", err.Error())
		os.Exit(1)
	}
	json.Indent(&buff, body, "", "\t")
	fmt.Println(string(buff.Bytes()))
}
