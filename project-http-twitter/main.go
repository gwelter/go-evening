package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var id int = 1

type tweetResponse struct {
	ID int `json:"ID"`
}

type tweetPayload struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

func tweets(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	t := tweetPayload{}

	if err := json.Unmarshal(body, &t); err != nil {
		log.Println("Failed tu unmarshal payload:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if t.Message == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("Tweet: `%s` from %s\n", t.Message, t.Location)
	payload, err := json.Marshal(tweetResponse{ID: id})
	if err != nil {
		log.Println("Failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(payload)
}

func main() {
	http.HandleFunc("/tweets", tweets)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
