package main

import (
	"log"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

// Function which performs a GET call
func getData(url string, ch chan<- result) {
	startTime := time.Now()
	if resp, error := http.Get(url); error != nil {
		ch <- result{
			url:     url,
			err:     error,
			latency: 0,
		}
	} else {
		ch <- result{
			url:     url,
			err:     nil,
			latency: time.Since(startTime).Round(time.Microsecond),
		}
		resp.Body.Close()
	}
}

func main() {
	ch := make(chan result)
	listOfUrls := [5]string{
		"https://www.google.com",
		"https://www.reddit.com",
		"hxxps://www.facebook.com",
		"https://www.nasdaq.com",
		"https://www.casio.com",
	}

	for _, v := range listOfUrls {
		go func(v string) {
			getData(v, ch)
		}(v)
	}

	for i := 0; i < len(listOfUrls); i++ {
		temp := <-ch
		if temp.err != nil {
			log.Printf("%s,%s", temp.url, temp.err)
		} else {
			log.Printf("%s,%s", temp.url, temp.latency)
		}
	}
}

// TODO implement done call
