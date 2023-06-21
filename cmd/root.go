package cmd

import (
	"fmt"
	"os"

	"github.com/jcardenasc93/gapi/handlers"
	"github.com/jcardenasc93/gapi/ui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gapi [URL]",
	Short: "gAPI is a CLI http client for developers",
	Long: `gAPI is a CLI http client for developers that allows to
    make http requests in an easy, interactive and straightforward way. Also allows
    workflows creation based on files. Enjoy it!`,
	Args: cobra.MinimumNArgs(1),
	Run:  runRequest,
}

func runRequest(cmd *cobra.Command, args []string) {
	url := args[0]
	resp, err := handlers.MakeRequest(url, "application/json", "GET")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pprinter := ui.NewPPrinter(resp)
	pprinter.PrintResp()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
