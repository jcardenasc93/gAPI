package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gapi [URL]",
	Short: "gAPI is a CLI http client for developers",
	Long: `gAPI is a CLI http client for developers that allows to
    make http requests in an easy, interactive and straightforward way. Also allows
    workflows creation based on files. Enjoy it!`,
	Run: runRequest,
}

func runRequest(cmd *cobra.Command, args []string) {
	a := strings.Join(args, " -- ")
	fmt.Printf("Run with %s", a)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
