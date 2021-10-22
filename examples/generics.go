package main

import (
	"fmt"
)

type Addable interface {
	type int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, complex64, complex128, string
}

func add[T Addable](values ...T) (sum T) { // HL
	for _, value := range values {
		sum += value
	}
	return
}

func main() {
	fmt.Println(add(1, 2)) // HL
	fmt.Println(add("foo", "bar", "baz")) // HL
	fmt.Println(add(1.1, 2.1, 3.1, 4.1, 5.1)) // HL
}
