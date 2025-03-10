package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func main() {

	// FileServer para servir la carpeta 'template' como estáticos
	fs := http.FileServer(http.Dir("static"))

	// Todo lo que venga por /static/ se va a servir desde la carpeta 'template'
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Ruta principal
	http.HandleFunc("/", Index)

	fmt.Println("El servidor está corriendo en http://localhost:3000")
	http.ListenAndServe("localhost:3000", nil)
}
