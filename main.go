package main

import "syscall/js"

func main() {
	channel := make(chan interface{}, 0)
	js.Global().Set("add", js.FuncOf(add))
	<-channel
}

func add(this js.Value, args []js.Value) interface{} {
	return args[0].Int() + args[1].Int()
}
