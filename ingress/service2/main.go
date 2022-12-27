package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	seed := rand.Intn(1000)

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[Service2] Request on node %d\n", seed)
		w.Write([]byte(fmt.Sprintf("Data from service2 node %d\n", seed)))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[Service2] Request on node %d path '%s'\n", seed, r.URL.Path)
		w.Write([]byte(fmt.Sprintf("[Service2] Ok on node %d\n", seed)))
	})

	fmt.Printf("Current node %v\n", seed)
	fmt.Printf("Server is up on port 1234\n")
	http.ListenAndServe(":1234", nil)
}
