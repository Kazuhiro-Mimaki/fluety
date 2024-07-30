package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	r := flag.Bool("r", false, "Receiver for centralized logging")
	s := flag.Bool("s", false, "Sender to centralized logging")
	flag.Parse()

	if *r {
		receiver := NewReceiver()

		fmt.Print("Start receiver")

		http.HandleFunc("/", receiver.Render())
		http.HandleFunc("/register", receiver.Register())
		http.HandleFunc("/read", receiver.Read())
		http.ListenAndServe(":8080", nil)
	}

	if *s {
		sender := NewSender()

		fmt.Printf("Start sender %v", sender.Id)

		in := bufio.NewScanner(os.Stdin)
		for in.Scan() {
			err := sender.Send(in.Text())
			if err != nil {
				panic(err)
			}
		}
	}
}
