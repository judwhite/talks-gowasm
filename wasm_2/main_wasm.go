package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"syscall/js"
	"time"

	"github.com/judwhite/talks-gowasm/wasm_2/wasm"
)

func main() {
	var (
		ul      *wasm.Element
		history *wasm.History
	)

	document := wasm.NewDOM()
	runButton := document.CreateElement("button")
	runButton.SetAttribute("id", "runButton")
	runButton.AppendChild(document.CreateTextNode("Run"))

	var wg sync.WaitGroup
	wg.Add(1)

	cb := js.NewCallback(func(args []js.Value) {
		start := time.Now()
		wg.Wait()
		fmt.Println("button click callback: started")

		node := document.CreateElement("li")
		textNode := document.CreateTextNode(document.Location())
		node.AppendChild(textNode)
		ul.AppendChild(node)

		length := len(document.GetElementsByTagName("li"))
		history.PushState(fmt.Sprintf("/history/%d", length))

		go func() {
			loc, _ := url.Parse(document.Location())
			getURL := fmt.Sprintf("http://%s:%s", loc.Hostname(), loc.Port())
			resp, _ := http.Get(getURL)
			defer resp.Body.Close()
			b, _ := ioutil.ReadAll(resp.Body)
			fmt.Printf("body length = %d, time taken %v\n", len(b), time.Since(start).Round(time.Millisecond))
		}()
		fmt.Printf("button click callback: finished, time taken %v\n", time.Since(start).Round(time.Millisecond))
	})
	runButton.Click(cb)

	app := document.GetElementById("app")
	go func() {
		defer wg.Done()
		history = wasm.NewHistory()

		app.AppendChild(runButton)

		ul = document.CreateElement("ul")
		app.AppendChild(ul)
	}()

	wasm.Wait() // HL

	fmt.Println("Bye Wasm!") // HL
}
