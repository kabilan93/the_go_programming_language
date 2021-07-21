package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex // create a mutext key
var count int

func main() {

	http.HandleFunc("/", handler) //each request calls handler
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	//fmt.Println("Initial setup")
}

//handler echoes the Path component of the requested URL.
// request is recieved as a struct
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Patch = %q\n", r.URL.Path)
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
