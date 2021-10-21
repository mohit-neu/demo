package main

import (
	"fmt"
	"log"
	"net/http"
)

//START OMIT
func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8001", nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	if path != "" {
		fmt.Fprintln(w, "Hello, %s!\n", path)
		return
	}

	fmt.Fprintf(w, "Hello, stranger!\n")
}

//END OMIT
