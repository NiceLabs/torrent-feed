package main

import (
	"embed"
	"flag"
	"io/fs"
	"log"
	"net/http"

	"github.com/NiceLabs/torrent-feed/internal/proxy"
)

var addr string

//go:embed files
var files embed.FS

func init() {
	flag.StringVar(&addr, "addr", ":8080", "http listen address")
	flag.Parse()
}

func main() {
	handler := &proxy.Handler{HTTPClient: http.DefaultClient}
	handler.FileBase, _ = fs.Sub(files, "files")
	log.Printf("Starting server on %q\n", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		panic(err)
	}
}
