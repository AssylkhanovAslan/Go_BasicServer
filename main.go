package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func initHandlers() {
	http.HandleFunc(HANDLER_INDEX, indexHandler)
	http.HandleFunc(HANDLER_FEEDBACK, feedbackHandler)
	http.HandleFunc(HANDLER_QUERY_GET, queryGet)
	http.HandleFunc(HANDLER_POST_JSON, postJson)
	http.HandleFunc(HANDLER_POST_FORM, postForm)
	http.HandleFunc(HANDLER_PATH, path)
}

func main() {
	initHandlers()

	err := http.ListenAndServe(":80", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}
