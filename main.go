package main

import (
	"awesomeProject1/list"
	"fmt"
)

func main() {
	l := list.NewList()
	l.Add(2)
	l.Add(12)
	l.Add(3)
	l.Add(1222)
	l.Add(1222)
	l.Print()

	fmt.Println()
	l.Print()
}
