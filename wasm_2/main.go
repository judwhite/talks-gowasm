package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var cache = make(map[string][]byte)

func main() {
	mustCache("wasm_exec.html")
	mustCache("wasm_exec.js")
	mustCache("encoding.js")
	mustCache("test.wasm.gz")
	mustCache("favicon.ico")

	http.HandleFunc("/favicon.ico", handleFavIcon)
	http.HandleFunc("/wasm_exec.js", handleWASMExec)
	http.HandleFunc("/encoding.js", handleEncodingJS)
	http.HandleFunc("/test.wasm", handleWASM)
	http.HandleFunc("/", handleIndex)

	addr := "127.0.0.1:8000"
	log.Printf("Listening on %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func mustCache(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	cache[filename] = b
}

func serveFile(w http.ResponseWriter, filename, contentType string) {
	w.Header().Set("content-type", contentType)

	if strings.HasSuffix(filename, ".gz") {
		w.Header().Set("content-encoding", "gzip")
	}

	w.Write(cache[filename])
}

func handleIndex(w http.ResponseWriter, _ *http.Request) {
	serveFile(w, "wasm_exec.html", "text/html")
}

func handleEncodingJS(w http.ResponseWriter, _ *http.Request) {
	serveFile(w, "encoding.js", "application/javascript")
}

func handleWASMExec(w http.ResponseWriter, _ *http.Request) {
	serveFile(w, "wasm_exec.js", "application/javascript")
}

func handleWASM(w http.ResponseWriter, _ *http.Request) {
	serveFile(w, "test.wasm.gz", "application/wasm")
}

func handleFavIcon(w http.ResponseWriter, _ *http.Request) {
	serveFile(w, "favicon.ico", "image/x-icon")
}
