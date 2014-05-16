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
		reader, err := r.MultipartReader()
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			return
		}

		for {
			part, err := reader.NextPart()
			if err != nil {
				break
			}
			defer part.Close()

			img, err := os.Create(part.FileName())
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(500)
				return
			}
			defer img.Close()

			io.Copy(img, part)
		}
	})

	m.Run()
}
