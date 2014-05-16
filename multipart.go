package main

import (
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"io"
	"net/http"
	"os"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Layout:     "layout",
		Extensions: []string{".html"}}))

	m.Get("/", func(r render.Render) {
		r.HTML(200, "index", nil)
	})

	m.Post("/upload", func(w http.ResponseWriter, r *http.Request) {
		key := "files[]"
		// reader, err := r.MultipartReader() // stream
		r.ParseMultipartForm(25)
		for _, fileHeader := range r.MultipartForm.File[key] {
			file, err := fileHeader.Open()
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(500)
				return
			}
			defer file.Close()
			img, err := os.Create(fileHeader.Filename)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(500)
				return
			}
			defer img.Close()
			io.Copy(img, file)
		}
	})

	m.Run()
}
