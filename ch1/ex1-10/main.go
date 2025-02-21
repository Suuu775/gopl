package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
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
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	defer resp.Body.Close() // don't leak resources

	// Save response to a file
	filename := fmt.Sprintf("response_%s_%d.txt", url, time.Now().UnixNano())
	file, err := os.Create(filename)
	if err != nil {
		ch <- fmt.Sprintf("while creating file %s: %v", filename, err)
		return
	}
	defer file.Close()

	nbytes, err := io.Copy(file, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s  saved to %s", secs, nbytes, url, filename)
}
