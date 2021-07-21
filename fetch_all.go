package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	//fmt.Println("Initial project setup")
	start := time.Now()

	ch := make(chan string)

	// starts go routine and send urls to the fetch func to make the requests
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err) // send to chanel ch
		return
	}

	//nBytes, err := io.Copy(ioutil.Discard, resp.Body)
	b, err := io.ReadAll(resp.Body)
	resp.Body.Close() // Dont leak resources

	if err != nil {
		ch <- fmt.Sprintf("while reading %s : %v", url, err)
		return
	}

	// write url output to file
	filename := strings.Split(url, "://")[1] + ".txt"
	fmt.Println("test", filename)
	wfErr := os.WriteFile("./"+filename, b, 0644)
	if wfErr != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, wfErr)
		return
	}

	secs := time.Since(start).Seconds()
	//ch <- fmt.Sprintf("%.2fs %7d %s", secs, nBytes, url) // write the output into the channel
	ch <- fmt.Sprintf("%.2fs %s \n %7d", secs, url, len(b)) // write the output into the channel

}
