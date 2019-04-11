package rest

import (
	"fmt"
	"net/http"
)

type Handler struct {
}

func (h *Handler) Routes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.handle(w, r)
	}
}

func (h *Handler) handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "rest.Handler.handle called")
}
