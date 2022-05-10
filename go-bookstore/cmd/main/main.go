package main

import (
	"log"
	"net/http"

	"github.com/faea726/go-bookstore/pkg/routers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routers.RegisterBookStoreRouter(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
