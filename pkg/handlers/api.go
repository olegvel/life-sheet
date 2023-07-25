package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/olegvel/life-sheet/pkg/models"
	"io"
	"log"
	"net/http"
)

func NewAPIServer() http.Handler {
	r := chi.NewRouter()
	h := &handler{}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	// RESTy routes for "media" resource
	r.Route("/media", func(r chi.Router) {
		r.Get("/", h.ListMedia)    // GET /media
		r.Post("/", h.CreateMedia) // POST /media

		// sub-routers:
		r.Route("/{mediaID}", func(r chi.Router) {
			r.Get("/", h.GetMedia)       // GET /media/123
			r.Put("/", h.UpdateMedia)    // PUT /media/123
			r.Delete("/", h.DeleteMedia) // DELETE /media/123
		})
	})

	return r
}

type handler struct {
	//TODO add usecase
}

func (h *handler) CreateMedia(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()

	rawBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	var item models.Item
	if err := json.Unmarshal(rawBody, &item); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	//_, err = h.usecase.CreateItem(ctx, item.ToModel())
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	log.Println(err)
	//	return
	//}

	w.WriteHeader(http.StatusCreated)
}

func (h *handler) ListMedia(w http.ResponseWriter, r *http.Request) {}

func (h *handler) GetMedia(w http.ResponseWriter, r *http.Request) {}

func (h *handler) UpdateMedia(w http.ResponseWriter, r *http.Request) {}

func (h *handler) DeleteMedia(w http.ResponseWriter, r *http.Request) {}
