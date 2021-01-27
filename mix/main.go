package main

import "fmt"
import "net/http"
import "html/template"
import "path"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// not yet implemented
		var filepath = path.Join("views", "index.html")
		var tmpl, err = template.ParseFiles(filepath)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var data = map[string]interface{}{
			"title": "Learning Golang Web",
			"name": "Batman",
		}

		err = tmpl.Execute(w, data)
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Server Started At localhost:9000")
	http.ListenAndServe(":9000", nil)
}