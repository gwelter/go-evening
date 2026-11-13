package main

import (
	"log"
	"my-twitter/server"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	s := server.Server{
		TweetReposity: &server.TweetsMemoryRepository{},
	}
	r.Get("/tweets", s.ListTweets)
	r.Post("/tweets", s.AddTweet)
	log.Fatal(http.ListenAndServe(":8080", r))
}
