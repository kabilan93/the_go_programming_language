package main

import (
	"fmt"
	"image/color"
	"log"
	"net/http"
)

var palette = []color.Color{color.White, color.Black, color.RGBA{0x00, 0xFF, 0x00, 0xFF}}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
	greenIndex = 2
)

func main() {

	http.HandleFunc("/", handler) //each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	//fmt.Println("Initial setup")
}

//handler echoes the Path component of the requested URL.
// request is recieved as a struct
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "HOST = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}

}
