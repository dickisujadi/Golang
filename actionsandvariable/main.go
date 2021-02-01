package main

import (
	"net/http"
	"fmt"
	"html/template"
)

type Info struct {
	Affiliation string
	Address string
}

type Person struct{
	Name string
	Gender string
	Hobbies []string
	Info Info
}



func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var person = Person {
			Name: "Bruce Wayne",
			Gender: "Male",
			Hobbies: []string{"Read Books", "Play Futsal", "Buy Things"},
			Info: Info{"Wayne Int", "Gotham City"},
		}

		var tmpl = template.Must(template.ParseFiles("view.html"))
		if err := tmpl.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}