package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type Resp struct {
	Message string `json:"message"`
	XResult string `json:"x-result"`
	XBody   string `json:"x-body"`
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	bodyBytes, _ := io.ReadAll(r.Body)
	_ = r.Body.Close()
	bodyStr := string(bodyBytes)

	xTest := r.Header.Get("x-test")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "x-test,ngrok-skip-browser-warning,Content-Type,Accept,Access-Control-Allow-Headers")

	resp := Resp{
		Message: "levchik",
		XResult: xTest,
		XBody:   bodyStr,
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		http.Error(w, `{"error":"encode failed"}`, http.StatusInternalServerError)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/result4/", resultHandler)

	http.HandleFunc("/result4", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/result4/", http.StatusMovedPermanently)
	})

	log.Printf("Server listening on :%s (endpoint: /result4/)", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}