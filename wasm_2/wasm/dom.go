package wasm

import "syscall/js"

type DOM struct {
	document js.Value
}

func NewDOM() *DOM {
	dom := &DOM{
		document: js.Global().Get("document"),
	}
	return dom
}

func (dom *DOM) CreateElement(el string) *Element {
	return &Element{dom.document.Call("createElement", el)}
}

func (dom *DOM) GetElementById(id string) *Element {
	return &Element{dom.document.Call("getElementById", id)}
}

func (dom *DOM) GetElementsByTagName(tag string) []*Element {
	value := dom.document.Call("getElementsByTagName", tag)
	var els []*Element
	for i := 0; i < value.Length(); i++ {
		els = append(els, &Element{value.Index(i)})
	}
	return els
}

func (dom *DOM) GetElementsByClassName(tag string) []*Element {
	value := dom.document.Call("getElementsByClassName", tag)
	var els []*Element
	for i := 0; i < value.Length(); i++ {
		els = append(els, &Element{value.Index(i)})
	}
	return els
}

func (dom *DOM) CreateTextNode(text string) *TextNode {
	return &TextNode{dom.document.Call("createTextNode", text)}
}

func (dom *DOM) Alert(text string) {
	dom.document.Call("alert", text)
}

func (dom *DOM) Location() string {
	return dom.document.Get("location").String()
}

func (dom *DOM) SetLocation(location string) {
	dom.document.Set("location", location)
}
