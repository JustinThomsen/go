package main

import "fmt"

func main() {
	greeter := gophersGreeter{"Hello", "gophers"}
	greeter.greet()
}

type gophersGreeter struct {
	how string
	who string
}

func (greeter gophersGreeter) greet() {
	fmt.Printf("%s %s!", greeter.how, greeter.who)
}
