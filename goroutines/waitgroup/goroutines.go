package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// START OMIT
func downloadURLs(urls ...string) {
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, u := range urls {
		go func(u string) {
			downloadUrl(u)
			wg.Done()
		}(u)
	}
	wg.Wait()
}

func main() {
	s := time.Now()
	downloadURLs("https://avocode.com", "https://golang.org", "https://google.com")
	fmt.Printf("Took %s\n", time.Now().Sub(s))
}

// END OMIT

func downloadUrl(url string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("url %s %d\n", url, len(b))
}
