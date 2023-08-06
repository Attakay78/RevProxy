package main

import (
	"net/http"
	"revproxy/server1"
)

func main(){
	srv := server1.NewServer()
	http.ListenAndServe(":8080", srv)
}
