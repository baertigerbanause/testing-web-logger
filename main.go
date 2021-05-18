package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func handlerReadyz(w http.ResponseWriter, r *http.Request) {
	log.Printf("[ %s ] [ %s ] : %s\n", r.RemoteAddr, r.Method, r.URL)
	fmt.Fprint(w, "OK\n", r.URL.Path[1:])
}
func handlerBase(w http.ResponseWriter, r *http.Request) {
	log.Printf("[ %s ] [ %s ] : %s\n", r.RemoteAddr, r.Method, r.URL)
	http.Redirect(w, r, "/health", http.StatusFound)
}

func main() {
	portPtr := flag.Int("port", 6443, "port to listen on")
	flag.Parse()
	port := ":" + strconv.Itoa(*portPtr)

	http.HandleFunc("/health", handlerReadyz)
	http.HandleFunc("/", handlerBase)
	log.Println("registered handlers")
	log.Println("Listening on", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
