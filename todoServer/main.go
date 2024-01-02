package main

import (
	"flag"
	"fmt"
	"net/http"

	"os"
	"time"
)

func main() {
	host := flag.String("h", "localhost", "Server Host")
	port := flag.Int("p", 8080, "Server Port")
	todoFile := flag.String("f", "todoServer.json", "todo JSON file")
	flag.Parse()

	mux := newMux(*todoFile)
	s := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", *host, *port),
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// testList := todo.List{}
	// testList.Get("./todo.json")
	// fmt.Println(testList)
	fmt.Printf("Server running on port %d", *port)
	if err := s.ListenAndServe(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
