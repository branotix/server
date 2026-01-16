package main

import (
	"fmt"
	"net/http"
)

func searchHanlder(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	limit := r.URL.Query().Get("limit")

	fmt.Fprintf(w, "you are search for%s\n: ", query)
	fmt.Fprintf(w, "and your limit is %s ", limit)
}

func main() {
	http.HandleFunc("/search", searchHanlder)
	fmt.Println("server lisitening: localhost:8080")
	http.ListenAndServe(":8080", nil)
}
