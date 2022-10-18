package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const QUERY_SEARCH = "search"

const HANDLER_INDEX = "/"
const HANDLER_FEEDBACK = "/feedback"
const HANDLER_QUERY_GET = "/query_get"
const HANDLER_POST_JSON = "/post_json"
const HANDLER_POST_FORM = "/post_form"
const HANDLER_PATH = "/path/"

func indexHandler(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("IndexHandler: Received request")
	io.WriteString(responseWriter, "I alive")
}

func feedbackHandler(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("FeedbackHandler: Received request")
	io.WriteString(responseWriter, "Thanks for the feedback")
}

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
	urlPath := request.URL.Path
	paramId := strings.TrimPrefix(urlPath, HANDLER_PATH)

	response := fmt.Sprintf("Thank you for path request. You request smth with id = %v\n", paramId)

	io.WriteString(responseWriter, response)
}
