package wasm

import "syscall/js"

type Element struct {
	js.Value
}

func (el *Element) AppendChild(node Value) {
	el.Call("appendChild", node.JSValue())
}

func (el *Element) AddEventListener(event string, callback js.Callback) {
	el.Call("addEventListener", event, callback)
}

func (el *Element) Click(callback js.Callback) {
	el.AddEventListener("click", callback)
}

func (el *Element) GetAttribute(name string) string {
	return el.Call("getAttribute", name).String()
}

func (el *Element) RemoveAttribute(name string) {
	el.Call("removeAttribute", name)
}

func (el *Element) SetAttribute(name, value string) {
	el.Call("setAttribute", name, value)
}

func (el *Element) HasAttribute(name string) bool {
	return el.Call("hasAttribute", name).Bool()
}

func (el *Element) JSValue() js.Value {
	return el.Value
}

