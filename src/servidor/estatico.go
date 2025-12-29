package main

import (
	"log"
	"net/http"
)

// função principal
func main() {
	fs := http.FileServer(http.Dir("./estatico"))
	http.Handle("/", fs)
	log.Println("Listening on: 3000...", nil)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
