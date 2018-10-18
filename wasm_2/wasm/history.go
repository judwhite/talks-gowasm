package wasm

import "syscall/js"

type History struct {
	js.Value
}

func NewHistory() *History {
	return &History{js.Global().Get("window").Get("history")}
}

func (h *History) PushState(url string) {
	h.Call("pushState", "{}", "", url)
}
