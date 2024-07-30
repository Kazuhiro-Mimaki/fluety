package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

const MAX_SENDERS_COUNT = 2

type Receiver struct {
	records    []Record
	registered map[string]*Sender
}

func NewReceiver() *Receiver {
	return &Receiver{
		records:    []Record{},
		registered: map[string]*Sender{},
	}
}

func (r *Receiver) Render() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		template, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := template.Execute(w, "no data needed"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (r *Receiver) Register() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		var record Record
		err := json.NewDecoder(req.Body).Decode(&record)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if _, ok := r.registered[record.Id]; !ok {
			r.registered[record.Id] = NewSender()
			return
		}
		if len(r.registered) > MAX_SENDERS_COUNT {
			http.Error(w, "Max senders reached", http.StatusInternalServerError)
			return
		}
		r.records = append(r.records, record)
	}
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
