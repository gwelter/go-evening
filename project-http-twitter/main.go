package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

type tweetsReposiory interface {
	AddTweet(t tweet) (int, error)
	Tweets() ([]tweet, error)
}

type tweetsMemoryRepository struct {
	tweets []tweet
}

func (t *tweetsMemoryRepository) AddTweet(tw tweet) (int, error) {
	t.tweets = append(t.tweets, tw)
	return len(t.tweets), nil
}

func (t tweetsMemoryRepository) Tweets() ([]tweet, error) {
	return t.tweets, nil
}

func (s server) addTweet(w http.ResponseWriter, r *http.Request) {
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

	id, err := s.tweetReposity.AddTweet(t)
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

func (s server) listTweets(w http.ResponseWriter, r *http.Request) {
	tweets, err := s.tweetReposity.Tweets()
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

func (s server) tweets(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		s.addTweet(w, r)
	} else if r.Method == http.MethodGet {
		s.listTweets(w, r)
	}
}

type server struct {
	tweetReposity tweetsReposiory
}

func main() {
	s := server{
		tweetReposity: &tweetsMemoryRepository{},
	}
	http.HandleFunc("/tweets", s.tweets)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
