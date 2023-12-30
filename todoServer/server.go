package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/handlers"
)

func newMux(todoFile string) http.Handler {
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"0.0.0.0/0"}), // Set your allowed origins here
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)

	m := http.NewServeMux()
	mu := &sync.Mutex{}

	m.HandleFunc("/", rootHandler)
	t := todoRouter(todoFile, mu)
	m.Handle("/todo", http.StripPrefix("/todo", t))
	m.Handle("/todo/", http.StripPrefix("/todo/", t))

	return corsHandler(m)
}

func replyTextContent(w http.ResponseWriter, r *http.Request, status int, content string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(status)
	w.Write([]byte(content))
}

func replyJSONContent(w http.ResponseWriter, r *http.Request, status int, resp *todoResponse) {
	body, err := json.Marshal(resp)
	if err != nil {
		replyError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(status)
	w.Write(body)
}

func replyError(w http.ResponseWriter, r *http.Request, status int, message string) {
	log.Printf("%s %s: Error: %d %s", r.URL, r.Method, status, message)
	http.Error(w, http.StatusText(status), status)
}
