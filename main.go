package main

import (
	"fmt"
	"io"
	"net/http"
)

func indexHandler(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("IndexHandler: Received request")
	fmt.Printf("Method: %v\n", request.Method)
	io.WriteString(responseWriter, "I alive")
}

func feedbackHandler(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("FeedbackHandler: Received request")
	fmt.Printf("Method: %v\n", request)
	io.WriteString(responseWriter, "Thanks for the feedback")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/feedback", feedbackHandler)

	err := http.ListenAndServe(":80", nil)

	fmt.Println(err)
}
