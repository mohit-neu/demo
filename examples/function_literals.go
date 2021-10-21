package main

import "fmt"

var f = func(x, y int) int { return x + y }

func main() {
	fmt.Println(f(10, 20))
}
