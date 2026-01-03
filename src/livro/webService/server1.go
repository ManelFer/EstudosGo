package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
