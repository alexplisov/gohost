package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	path := flag.String("path", "", "Specify directory path to host on local server.")
	address := flag.String("address", "localhost:3000", "Specify address to run local server on.")
	flag.Parse()

	if *path == "" {
		fmt.Println("No directory path specified.", "Specify directory path to host on local server.")
		return
	}

	fsys := os.DirFS(*path)
	handler := http.FileServerFS(fsys)

	fmt.Printf("Local server for directory %s started on %s\n", *path, *address)

	http.ListenAndServe(*address, handler)
}
