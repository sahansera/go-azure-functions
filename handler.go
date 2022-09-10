package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
)

var quotes = []string{
	"Talk is cheap. Show me the code.",
	"First, solve the problem. Then, write the code.",
	"Experience is the name everyone gives to their mistakes.",
	"Any fool can write code that a computer can understand. Good programmers write code that humans can understand.",
}

func quotesHandler(w http.ResponseWriter, r *http.Request) {
	// Get a random quote
	message := quotes[rand.Intn(len(quotes))]

	// Write the response
	fmt.Fprint(w, message)
}

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	http.HandleFunc("/api/GetQuotes", quotesHandler)
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
