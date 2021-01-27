package main

import "fmt"
import "net/http"

func main() {
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	fmt.Println("Server started at localhost:90000")
	http.ListenAndServe(":9000", nil)
}