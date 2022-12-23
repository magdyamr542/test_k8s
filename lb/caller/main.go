package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	fmt.Printf("Caller started\n")
	counter := 0
	// Create the http client that disables connection pooling (to test load balancing)
	httpTransport := http.DefaultTransport.(*http.Transport).Clone()
	httpTransport.DisableKeepAlives = true
	httpClient := &http.Client{Transport: httpTransport}

	for {
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("[%d] Doing a GET request to /hello\n", counter)
		resp, err := httpClient.Get("http://lb:4000/api/v1/hello")
		counter += 1
		if err != nil {
			fmt.Printf("Error with GET request to /hello %v\n", err.Error())
			continue
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error with reading body %v\n", err.Error())
			continue
		}
		//Convert the body to type string
		data := string(body)
		fmt.Printf("%v\n", data)
		err = resp.Body.Close()
		if err != nil {
			fmt.Printf("Error while closing the body %v\n", err.Error())
		}
	}
}
