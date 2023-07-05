package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/jcardenasc93/gapi/internal/handler"
	"github.com/jcardenasc93/gapi/internal/ui"
	"github.com/spf13/cobra"
)

var (
	headersFlag string
)

var rootCmd = &cobra.Command{
	Use:   "gapi [method] URL",
	Short: "gAPI is a CLI http client for developers",
	Long: `gAPI is a CLI http client for developers that allows to
    make http requests in an easy, interactive and straightforward way. Also allows
    workflows creation based on files. Enjoy it!`,
	Args: cobra.MinimumNArgs(1),
	Run:  runRequest,
}

func runRequest(cmd *cobra.Command, args []string) {
	var (
		method string
		url    string
	)

	if len(args) == 1 {
		method = handler.DefaultVerb
		url = args[0]
	} else {
		method = args[0]
		url = args[1]
	}
	hs := strings.Split(headersFlag, "|")
	h, err := handler.MakeReq(url, method, hs)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		os.Exit(1)
	}
	pp := ui.NewPPrinter(h.Req, h.Resp)
	pp.Print()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&headersFlag, "headers", "H", "", `
        Add pipe separated headers as key:value.
        Example: Content-Type:application/json|Accept-Language:en-US,en
        `)
}
