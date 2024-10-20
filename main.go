package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {
	fmt.Print("Start fluety")

	recorder := NewRecorder()

	go ScanStdin(&recorder)

	http.HandleFunc("/", RenderTemplate())
	http.HandleFunc("/read", SSEResponse(&recorder))
	http.ListenAndServe(":8080", nil)
}

func ScanStdin(recorder *Recorder) {
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		recorder.Enqueue(
			Record{
				Body: in.Text(),
			})
	}
}

func RenderTemplate() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		template, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if err := template.Execute(w, "no data needed"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func SSEResponse(recorder *Recorder) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Content-Type", "text/event-stream")

		for recorder.Exists() {
			head := recorder.Dequeue()
			jsonHead, err := json.Marshal(head)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			fmt.Fprintf(w, "data: %s \n\n", jsonHead)
			w.(http.Flusher).Flush()
		}
	}
}
