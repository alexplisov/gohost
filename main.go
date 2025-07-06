package main

import (
	"flag"
	"net/http"
	"os"
)

func main() {
	path := flag.String("path", "", "Specify directory path to host on local server.")
	address := flag.String("address", "localhost:3000", "Specify address to run local server on.")
	flag.Parse()

	fsys := os.DirFS(*path)
	handler := http.FileServerFS(fsys)
	http.ListenAndServe(*address, handler)
}
