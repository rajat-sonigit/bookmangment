package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rajat-sonigit/go-bookstore/pkg/routes"
)
func main() {
	r := mux.NewRouter()
	routes.RegisterBookStore(r)
	fmt.Printf("Port started")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
