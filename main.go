package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Print("Start fluety")

	receiver := NewReceiver()
	template := NewTemplate()

	go func() {
		in := bufio.NewScanner(os.Stdin)
		for in.Scan() {
			receiver.Add(in.Text())
		}
	}()

	http.HandleFunc("/", template.Render())
	http.HandleFunc("/read", receiver.Read())
	http.ListenAndServe(":8080", nil)
}
