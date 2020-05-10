package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("vim-go")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
