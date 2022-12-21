package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	serverVersion := os.Getenv("SERVER_VERSION")
	apiPrefix := os.Getenv("API_PREFIX")

	rand.Seed(time.Now().UnixNano())
	seed := rand.Intn(1000)

	with := func(apiPath string) string {
		apiPrefix = strings.TrimSuffix(apiPrefix, "/")
		apiPath = strings.TrimPrefix(apiPath, "/")
		return fmt.Sprintf("%s/%s", apiPrefix, apiPath)
	}

	http.HandleFunc(with("hello"), func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request on node %d\n", seed)
		w.Write([]byte(fmt.Sprintf("Hello from %d\n", seed)))
	})

	http.HandleFunc(with("/server"), func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request on node %d\n", seed)
		w.Write([]byte(fmt.Sprintf("Server version %s\n", serverVersion)))
	})

	http.HandleFunc(with("/dynamic"), func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request on node %d\n", seed)
		dynamicValue := os.Getenv("DYNAMIC")
		w.Write([]byte(fmt.Sprintf("Dynamic value %s\n", dynamicValue)))
	})

	fmt.Printf("Server is up on port 1234\n")
	fmt.Printf("Api accessible under '<host>:<port>%s'\n", apiPrefix)
	http.ListenAndServe(":1234", nil)
}
