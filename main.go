package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

const CAPACITY = 100

func main() {
	log.Println("Starting fluety")

	recorder := NewRecorder(CAPACITY)

	go scanning(recorder, os.Stdin)

	http.HandleFunc("/", renderTemplate)
	http.HandleFunc("/sse", streamRead(recorder))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func scanning(recorder *Recorder, r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		recorder.Enqueue(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error scanning input: %v", err)
	}
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func streamRead(recorder *Recorder) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		setupSSEHeaders(w)

		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
			return
		}

		for {
			record, exists := recorder.Dequeue()
			if !exists {
				break
			}
			jsonRecord, err := json.Marshal(record)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Fprintf(w, "data: %s\n\n", jsonRecord)
			flusher.Flush()
		}
	}
}

func setupSSEHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Content-Type", "text/event-stream")
}
