package main

import "fmt"

type Whistle string

func AcceptAnything(thing interface{}) {
	fmt.Println(thing)
	whistle, ok := thing.(Whistle)
	if ok {
		whistle.MakeSound()
	}
}
func AcceptWhistle(whistle Whistle) {
	fmt.Println(whistle)
	whistle.MakeSound()
}
func main() {
	AcceptAnything(3.1415)
	AcceptAnything("A string")
	AcceptAnything(true)
	AcceptAnything(Whistle("Toyco Canary"))
}
