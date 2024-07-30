package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type Sender struct {
	Id string
}

func NewSender() *Sender {
	return &Sender{
		Id: uuid.New().String(),
	}
}

func (s *Sender) Send(text string) error {
	record := Record{
		Id:   s.Id,
		Body: text,
	}
	jsonData, err := json.Marshal(record)
	if err != nil {
		return err
	}
	resp, err := http.Post("http://localhost:8080/register", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	err = resp.Body.Close()
	if err != nil {
		return err
	}
	return nil
}
