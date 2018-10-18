package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/wasm_exec.js", handleWASMExec)
	http.HandleFunc("/test.wasm", handleWASM)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func serveFile(w http.ResponseWriter, r *http.Request, filename, contentType string) {
	f, err := os.Open(filename)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	defer f.Close()

	w.Header().Set("content-type", contentType)
	b, err := ioutil.ReadAll(f)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(b)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	serveFile(w, r, "wasm_exec.html", "text/html")
}

func handleWASMExec(w http.ResponseWriter, r *http.Request) {
	serveFile(w, r, "wasm_exec.js", "application/javascript")
}

func handleWASM(w http.ResponseWriter, r *http.Request) {
	serveFile(w, r, "test.wasm", "application/wasm")
}
