package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	time.Sleep(400 * time.Millisecond)
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "[web2] Hello World from Web 2 (slow)! ")
	fmt.Fprintf(w, hostname)
	fmt.Fprintf(w, "\n")
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server started.")
	http.ListenAndServe(":80", nil)
}
