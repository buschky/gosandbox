
package hello

import (
//    "fmt"
    "net/http"
    "html/template"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
//    	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
//	fmt.Fprint(w, "Hello, world!")
//}

//func main() {
	tmpl := template.Must(template.ParseFiles("layout.html"))

//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
//	})

//	http.ListenAndServe(":80", nil)
}
