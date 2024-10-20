package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		// Initialize Receiver
		receiver := NewReceiver()

		fmt.Print("Start receiver")

		http.HandleFunc("/", receiver.Render())
		http.HandleFunc("/register", receiver.Register())
		http.HandleFunc("/read", receiver.Read())

		http.ListenAndServe(":8080", nil)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// Initialize Sender
		sender := NewSender()

		fmt.Printf("Start sender %v", sender.Id)

		in := bufio.NewScanner(os.Stdin)
		for in.Scan() {
			err := sender.Send(in.Text())
			if err != nil {
				panic(err)
			}
		}
	}()

	wg.Wait()
}
