package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Receiver struct {
	records []Record
}

func NewReceiver() *Receiver {
	return &Receiver{
		records: []Record{},
	}
}

func (r *Receiver) Add(text string) {
	record := Record{Body: text}
	r.records = append(r.records, record)
}

func (r *Receiver) Read() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Content-Type", "text/event-stream")

		for len(r.records) > 0 {
			head, rest := r.records[0], r.records[1:]
			r.records = rest
			jsonHead, err := json.Marshal(head)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Fprintf(w, "data: %s \n\n", jsonHead)
			w.(http.Flusher).Flush()
		}
	}
}
