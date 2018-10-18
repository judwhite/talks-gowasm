package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	var cb js.Callback
	cb = js.NewCallback(func(args []js.Value) {
		fmt.Println("button clicked")
		//cb.Release() // release the callback if the button will not be clicked again
	})

	runButton := js.Global().Get("document").Call("getElementById", "runButton")
	runButton.Call("addEventListener", "click", cb)

	select {} // wait forever. there's a better way, this will lead to out of memory.
}
