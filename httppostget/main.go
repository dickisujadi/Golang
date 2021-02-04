package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST" :
			w.Write([]byte("post"))
		case "GET" :
			w.Write([]byte("get"))
		default:
			http.Error(w, "", http.StatusBadRequest)
		}
	})

	fmt.Println("Server Started at Localhost:9000")
	http.ListenAndServe(":9000", nil)
}