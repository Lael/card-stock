package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func getIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

func ipHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello, %s! Your IP address is %s.", r.URL.Path[1:], getIP(r))
}

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", ipHandler)
	log.Printf("Listening at %s:%s!", hostname, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
