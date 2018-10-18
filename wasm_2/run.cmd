set GOOS=windows
set GOARCH=amd64
go build -o wasm.exe main.go

set GOOS=js
set GOARCH=wasm
go build -o test.wasm main_wasm.go
gzip -9 -f test.wasm

wasm
