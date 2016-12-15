package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":3001", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, 你好!"))
}
