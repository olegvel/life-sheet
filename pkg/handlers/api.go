package handlers

import (
	"encoding/json"
	"github.com/olegvel/life-sheet/pkg/models"
	"io"
	"log"
	"net/http"
)

func NewServer(port string) *http.Server {
	mux := http.NewServeMux()
	h := &handler{}

	fs := http.FileServer(http.Dir("./web"))
	mux.Handle("/", fs)

	// TODO add group for api
	mux.HandleFunc("/api/create", h.CreateHandler)

	return &http.Server{
		Addr:    port,
		Handler: mux,
	}

}

type handler struct {
	//TODO add usecase
}

func (h *handler) CreateHandler(w http.ResponseWriter, r *http.Request) {
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
