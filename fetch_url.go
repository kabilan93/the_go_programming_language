package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	cmdline_arg_url := os.Args[1:] // get access to arguments passed i cmd line when running the project

	for _, url := range cmdline_arg_url {
		//fmt.Println(item)
		//1.8

		prefix := "http://"

		if strings.HasPrefix(url, prefix) == false {
			url = prefix + url
		}

		resp, err := http.Get(url)
		// handle error with http get fuction
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//1.7
		body, err := io.Copy(os.Stdout, resp.Body)

		//1.9
		code := resp.StatusCode

		// body, err := ioutil.ReadAll(resp.Body)
		// resp.Body.Close()
		//handle error with reading response body
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Println(code)
		fmt.Println("\n", string(rune(body)))
	}
}
