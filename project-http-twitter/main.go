package main

import (
	"log"
	"my-twitter/server"
	"net/http"
)

func main() {
	s := server.Server{
		TweetReposity: &server.TweetsMemoryRepository{},
	}
	http.HandleFunc("/tweets", s.Tweets)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
