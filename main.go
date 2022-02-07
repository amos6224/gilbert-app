package main

import (
       "net/http"
        "fmt"
       "log"
       )

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/"{
		http.NotFound(w, r)
    }
	  fmt.Fprintf(w, `al konyen bouzin, Jeff genyen yo gwo zozo`)
	  return
  }

  func main() {
        mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	log.Println("Starting server.....")
	http.ListenAndServe(":4000", mux)
}
