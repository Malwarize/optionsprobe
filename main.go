package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	var concurrency int
	flag.IntVar(&concurrency, "c", 20, "max concurrency")
	var timeout int
	flag.IntVar(&timeout, "t", 100000, "timeout in millisecond")

	var headers string
	flag.StringVar(&headers, "H", "", "headers")

	flag.Parse()

	maxChan := make(chan struct{}, concurrency)
	scanner := bufio.NewScanner(os.Stdin)
	wg := sync.WaitGroup{}

	p := NewProbber("OPTIONS", time.Duration(timeout)*time.Millisecond, headers)

	for scanner.Scan() {
		maxChan <- struct{}{}
		url := scanner.Text()
		if url[:4] != "http" || url[:5] != "https" {
			url = "https://" + url
		}
		go func(u string) {
			defer func() { <-maxChan }()
			defer wg.Done()
			wg.Add(1)
			p.OptionProbe(u)
		}(url)
	}
	defer close(maxChan)
	wg.Wait()
}

type Probber struct {
	method  string
	headers string
	timeout time.Duration
	client  *http.Client
}

func NewProbber(method string, timeout time.Duration, headers string) *Probber {
	return &Probber{
		method:  method,
		timeout: timeout,
		headers: headers,
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (p *Probber) OptionProbe(url string) {
	optionRequest, _ := http.NewRequest(p.method, url, nil)
	for _, header := range strings.Split(p.headers, ",") {
		header = strings.TrimSpace(header)
		if header != "" {
			headerName := header[:strings.Index(header, ":")]
			headerValue := header[strings.Index(header, ":")+1:]
			optionRequest.Header.Set(headerName, headerValue)
		}
	}

	resp, err := p.client.Do(optionRequest)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if err != nil {
		return
	}
	go p.printMethods(url, resp)
}

func (p *Probber) printMethods(url string, resp *http.Response) {
	if resp.Header.Get("Allow") == "" {
		return
	}
	for _, method := range strings.Split(resp.Header.Get("Allow"), ",") {
		fmt.Print(url + " " + method + "\n")
	}
}
