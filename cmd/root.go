package cmd

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var port uint16

var rootCmd = &cobra.Command{
	Use:   "gohost <path>",
	Short: "Start a local static file server",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		fsys := os.DirFS(path)
		handler := http.FileServerFS(fsys)

		fmt.Printf("Local server for directory %s started on http://localhost:%d/\n", path, port)

		if err := http.ListenAndServe(net.JoinHostPort("localhost", strconv.FormatUint(uint64(port), 10)), handler); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().Uint16Var(&port, "port", 1337, "Optional. Port to listen on")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
