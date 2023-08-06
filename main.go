package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"revproxy/server2"
)

func main() {
	server2Router := mux.NewRouter()
	fmt.Println("Starting server at port 8000")
	server2.SetUpHandlers(server2Router)
	log.Fatal(http.ListenAndServe(":8000", server2Router))
}
