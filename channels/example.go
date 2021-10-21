package main

import "fmt"

func main() {
	messages := make(chan string, 10)

	for i := 0; i < 10; i++ {
		go func(x int) { messages <- fmt.Sprintf("ping %d", x) }(i)
	}

	for i := 0; i < 10; i++ {
		select {
		case msg := <-messages:
			fmt.Println(msg)
		}
	}
}
