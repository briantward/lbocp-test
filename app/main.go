package main

import "fmt"
import "net/http"
import "os"
import "log"

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	fmt.Fprintf(w, "Hi there, you have hit %s", hostname)
	// TODO get client identifer by cookie set by haproxy routers and IP

	clientname := "TODO"
	fmt.Printf("%s received a request from %s", hostname, clientname)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
