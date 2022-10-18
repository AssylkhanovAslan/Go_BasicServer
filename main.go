package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func indexHandler(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("IndexHandler: Received request")
	io.WriteString(responseWriter, "I alive")
}

func feedbackHandler(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("FeedbackHandler: Received request")
	io.WriteString(responseWriter, "Thanks for the feedback")
}

const QUERY_SEARCH = "search"

func queryGet(responseWriter http.ResponseWriter, request *http.Request) {
	searchQuery := request.URL.Query().Get(QUERY_SEARCH)
	fmt.Printf("SearchQuery = %v\n", searchQuery)
	if searchQuery == "" {
		responseWriter.WriteHeader(http.StatusBadRequest)
		io.WriteString(responseWriter, "{\"msg\":\"missing_field\"}")
		return
	}

	_, err := io.WriteString(responseWriter, "QueryGet request was processed. Thank you")
	if err != nil {
		return
	}
}

func postJson(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("PostJson. Received request")
	if request.Method != "POST" {
		io.WriteString(responseWriter, "Only waiting for POST request")
		return
	}

	decoder := json.NewDecoder(request.Body)
	var data post_data
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Printf("Error parsing data = %v\n", err)
	}

	fmt.Printf("Body = %v\n", data.Msg)
	io.WriteString(responseWriter, "Thank you for POST json request")
}

func postForm(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("PostForm. Received request")
	msg := request.FormValue("msg")

	fmt.Printf("Msg received = %v\n", msg)

	io.WriteString(responseWriter, "Thank you for POST form request")
}

func path(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Printf("url = %v\n", request.URL)
	io.WriteString(responseWriter, "Thank you for path request")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/feedback", feedbackHandler)
	http.HandleFunc("/query_get", queryGet)
	http.HandleFunc("/post_json", postJson)
	http.HandleFunc("/post_form", postForm)
	http.HandleFunc("/path/", path)

	err := http.ListenAndServe(":80", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}
