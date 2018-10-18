package wasm

import "syscall/js"

type TextNode struct {
	js.Value
}

func (t *TextNode) JSValue() js.Value {
	return t.Value
}
