package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Payload string `json:"payload"`
}

func main() {
	http.HandleFunc("/getMe", getMeHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func getMeHandler(w http.ResponseWriter, r *http.Request) {
	payload := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyLguYDguJ7guLfguYjguK3guJnguKrguJnguLTguJciOiLguIHguK3guILguYnguLLguKciLCLguYDguIjguLLguLDguKvguLkiOnRydWUsIuC4p-C4tOC4iuC4suC4l-C4teC5iOC4luC4meC4seC4lCI6IkNvbSBwcm8gYW5kIEFkdiBjb20gcHJvIn0.ewOyd-eUnFRxe9rt8JQfmY8JU3KFJxJRLf5kpWnEV0M"

	response := Response{
		Payload: payload,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
