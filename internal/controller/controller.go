package controller

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"html/template"
	"net/http"
	"test/internal/store"
)

type message struct {
	Message string `json:"message"`
}

func Build(r *chi.Mux, s store.Store) {
	r.Use(middleware.Logger)

	r.Get("/orders/{orderUid}", func(w http.ResponseWriter, r *http.Request) {
		getOrder(w, r, s)
	})
}

func getOrder(w http.ResponseWriter, r *http.Request, s store.Store) {
	orderUid := chi.URLParam(r, "orderUid")

	order, err := s.FindByOrderUID(orderUid)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	ts, err := template.ParseFiles("ui/html/order.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := ts.Execute(w, order); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
