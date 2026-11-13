package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type tweetResponse struct {
	ID int `json:"ID"`
}

type tweet struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

type tweetsList struct {
	Tweets []tweet `json:"tweets"`
}

func (s Server) addTweet(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	t := tweet{}

	if err := json.Unmarshal(body, &t); err != nil {
		log.Println("Failed tu unmarshal payload:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if t.Message == "" {
		log.Println("Missing message:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := s.TweetReposity.AddTweet(t)
	if err != nil {
		log.Println("Failed to add tweet:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Printf("Tweet: `%s` from %s\n", t.Message, t.Location)
	payload, err := json.Marshal(tweetResponse{ID: id})
	if err != nil {
		log.Println("Failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(payload)
}

func (s Server) listTweets(w http.ResponseWriter, r *http.Request) {
	tweets, err := s.TweetReposity.Tweets()
	if err != nil {
		log.Println("Failed get tweets:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	payload, err := json.Marshal(tweetsList{tweets})
	if err != nil {
		log.Println("Failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(payload)
}

func (s Server) Tweets(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		duration := time.Since(start)
		fmt.Printf("%s %s %s\n", r.Method, r.URL, duration)
	}()

	if r.Method == http.MethodPost {
		s.addTweet(w, r)
	} else if r.Method == http.MethodGet {
		s.listTweets(w, r)
	}
}

type Server struct {
	TweetReposity tweetsReposiory
}
