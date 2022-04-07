package main

import (
	"syscall/js"
)

func main() {
	channel := make(chan interface{}, 1)
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("createObject", js.FuncOf(createObject))
	js.Global().Set("createPerson", js.FuncOf(createPerson))
	<-channel
}

func add(this js.Value, args []js.Value) interface{} {
	return args[0].Int() + args[1].Int()
}

func createObject(this js.Value, args []js.Value) interface{} {
	return map[string]interface{}{
		"name": "Hong Gildong",
		"age":  26,
	}
}

func createPerson(this js.Value, args []js.Value) interface{} {
	person := js.Global().Call("Person", "Hong Cheolsoo", 26)
	return person
}
