/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the root command
var rootCmd = &cobra.Command{
	Use:   "gohost <path_to_folder>",
	Short: "Start a local static file server",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		// FIXME: change to dynamic
		port := 1337

		fsys := os.DirFS(path)
		handler := http.FileServerFS(fsys)

		fmt.Printf("Local server for directory %s started on http://localhost:%d/\n", path, port)

		http.ListenAndServe("localhost:1337", handler)
	},
}

func init() {
	// TODO: add port flag
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
