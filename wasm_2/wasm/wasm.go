package wasm

import "syscall/js"

var chanBeforeUnload = make(chan struct{})

func init() {
	beforeUnloadCb := js.NewEventCallback(0, func(_ js.Value) {
		close(chanBeforeUnload)
	})
	js.Global().Get("addEventListener").Invoke("beforeunload", beforeUnloadCb)
}

func Wait() {
	<-chanBeforeUnload
}
