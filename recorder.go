package main

type Record struct {
	Body string
}

type Recorder struct {
	records []Record
}

func NewRecorder() Recorder {
	return Recorder{
		records: []Record{},
	}
}

func (r *Recorder) Exists() bool {
	return len(r.records) != 0
}

func (r *Recorder) Enqueue(record Record) {
	r.records = append(r.records, record)
}

func (r *Recorder) Dequeue() Record {
	head, rest := r.records[0], r.records[1:]
	r.records = rest
	return head
}
