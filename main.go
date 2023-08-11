package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func setOriginServerURL(r *http.Request, originURL *url.URL){
	r.Host = originURL.Host
	r.URL.Host = originURL.Host
	r.URL.Scheme = originURL.Scheme
	r.RequestURI = ""
}

func main() {
	port := ":8080"
	originServer1, error1 := url.Parse("http://localhost:8081")

	if error1 != nil{
		panic(error1)
	}

	originServer2, error2 := url.Parse("http://localhost:8082")

	if error2 != nil{
		panic(error2)
	}

	reverseProxy := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[Reverse proxy server] received request from URL")

		if strings.HasPrefix(r.URL.String(), "/server1"){
			setOriginServerURL(r, originServer1)
		}else if strings.HasPrefix(r.URL.String(), "/server2"){
			setOriginServerURL(r, originServer2)
		}

		response, error := http.DefaultClient.Do(r)

		if error != nil{
			panic(error)
		}

		io.Copy(w, response.Body)

	})

	log.Fatal(http.ListenAndServe(port, reverseProxy))

}
