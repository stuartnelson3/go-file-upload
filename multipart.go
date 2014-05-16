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
		file, _, err := r.FormFile("files[]")
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			return
		}
		img, err := os.Create("img.jpg")
		defer img.Close()

		io.Copy(img, file)
	})

	m.Run()
}
