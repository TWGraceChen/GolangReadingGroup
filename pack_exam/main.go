package main

import (
	"fmt"
	"mypkg"
)

func main() {
	a := 10
	b := mypkg.Addone(a)
	fmt.Printf("%v + 1 = %v\n",a,b)
}