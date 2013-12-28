package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", &myHandler)
	mux.HandleFunc("/hello", sayHello)

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir(wd))))
}

type myHandler struct{}

func (*myHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "URL: "+r.URL.String())

}
func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World,this is version 2")
}
