package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler) //each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	//fmt.Println("Initial setup")
}

//handler echoes the Path component of the requested URL.
// request is recieved as a struct
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Patch = %q\n", r.URL.Path)
}
