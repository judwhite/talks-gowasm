set GOOS=js
set GOARCH=wasm
go build -o test.wasm main_wasm.go

set GOOS=windows
set GOARCH=amd64
@echo Open http://127.0.0.1:8080/wasm_exec.html
goexec "http.ListenAndServe(\"127.0.0.1:8080\", http.FileServer(http.Dir(\".\")))"
