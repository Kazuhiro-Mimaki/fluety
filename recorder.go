package main

import (
	"time"
)

type Record struct {
	Body      string
	Timestamp time.Time
}

type Recorder struct {
	records chan Record
}

func NewRecorder(capacity int) *Recorder {
	if capacity <= 0 {
		capacity = 100 // Default capacity if invalid value is provided
	}
	return &Recorder{
		records: make(chan Record, capacity),
	}
}

func (r *Recorder) Enqueue(body string) {
	record := Record{
		Body:      body,
		Timestamp: time.Now(),
	}
	select {
	case r.records <- record:
		// Record added successfully
	default:
		// Channel is full, log this event
		// You might want to implement a proper logging mechanism here
		println("Warning: Record channel is full. Dropping new record.")
	}
}

func (r *Recorder) Dequeue() (Record, bool) {
	select {
	case record := <-r.records:
		return record, true
	default:
		return Record{}, false
	}
}
