package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	layout         string = "2006-01-02 15:04:05" // time format
	responseString string = "Test body string."
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		userAgent := r.UserAgent()
		spaceIndex := strings.IndexByte(userAgent, ' ')
		if spaceIndex > 0 {
			userAgent = userAgent[:spaceIndex]
		}
		fmt.Printf("[%s] get request from %s; \"%s\" sent to requester\n", time.Now().Format(layout), userAgent, responseString)
		fmt.Fprint(w, responseString)
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
