package wasm

import "syscall/js"

type Value interface {
	JSValue() js.Value
}

