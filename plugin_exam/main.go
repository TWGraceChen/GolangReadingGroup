package main

import (
	"fmt"
	"plugin"
)

func main() {

	p, err := plugin.Open("plugin.so")
	if err != nil {
		panic(err)
	}
	m, err := p.Lookup("Addone")
	if err != nil {
		panic(err)
	}
	a := 10
	b := m.(func(int) int)(a)
	fmt.Printf("%v + 1 = %v\n",a,b)
}