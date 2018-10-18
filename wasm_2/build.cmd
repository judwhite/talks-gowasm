set GOOS=js
set GOARCH=wasm
go build -o test.wasm main_wasm.go
gzip -9 -f test.wasm
