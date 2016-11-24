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
		go func(url string) {
			defer wg.Done()
			lengthCh := make(chan int)
			go downloadUrl(url, lengthCh)

			timeout := time.After(300 * time.Millisecond)
			for {
				select {
				case length := <-lengthCh:
					fmt.Printf("URL %s %d\n", url, length)
					return
				case <-timeout:
					fmt.Printf("Downloading %s too long\n", url)
					return
				}
			}
		}(u)
	}
	wg.Wait()
}

// END OMIT

func main() {
	s := time.Now()
	downloadURLs("https://avocode.com", "https://golang.org", "https://google.com")
	fmt.Printf("Took %s\n", time.Now().Sub(s))
}

func downloadUrl(url string, c chan int) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	c <- len(b)
}
