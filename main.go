package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP HTTP リクエスト処理
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("templates",
				t.filename)))
	})

	if err := t.templ.Execute(w, nil); err != nil {
		log.Fatal("Template Execute:", err)
	}
}

func main() {
	// port
	const PORT string = ":8080"

	// route
	http.Handle("/", &templateHandler{filename: "index.html"})

	// web サーバーを開始
	fmt.Println("Server Runing: http://localhost" + PORT)
	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
