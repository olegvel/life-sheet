package handlers

import (
	"github.com/go-chi/chi/v5"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Name string
}

var indexFilePath = "./web/index.html"

func NewWebServer() http.Handler {
	r := chi.NewRouter()

	tmpl := template.Must(template.New("index.html").ParseFiles(indexFilePath))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, User{Name: "User Name"})
		if err != nil {
			log.Println(err)
			return
		}
	})

	r.Mount("/static", http.StripPrefix("/static", http.FileServer(http.Dir("./web/static"))))

	//// RESTy routes for "media" resource
	//r.Route("/media", func(r chi.Router) {
	//	r.Get("/", h.ListMedia)    // GET /media
	//	r.Post("/", h.CreateMedia) // POST /media
	//
	//	// sub-routers:
	//	r.Route("/{mediaID}", func(r chi.Router) {
	//		r.Get("/", h.GetMedia)       // GET /media/123
	//		r.Put("/", h.UpdateMedia)    // PUT /media/123
	//		r.Delete("/", h.DeleteMedia) // DELETE /media/123
	//	})
	//})

	return r
}
