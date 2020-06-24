package webfinger

import (
	"encoding/json"
	"errors"
	"net/http"
)

// HandlerOption handler option func
type HandlerOption func(h *Handler)

// Handler is a HTTP handler that implements the webinger protocol
type Handler struct {
	db          DB
	allowOrigin *string
}

// WithAllowOrigin sets Access-Control-Allow-Origin header to "orgn"
func WithAllowOrigin(orgn string) HandlerOption {
	return func(h *Handler) {
		h.allowOrigin = new(string)
		*h.allowOrigin = orgn
	}
}

// NewHandler returns a new handler instance
func NewHandler(db DB, opts ...HandlerOption) *Handler {

	h := &Handler{
		db: db,
	}

	for _, o := range opts {
		o(h)
	}

	return h
}

// ServeHTTP handles HTTP requests
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// parse user query
	q := QueryFromValues(r.URL.Query())

	// get ther resource from DB
	res, err := h.db.Get(q)
	if err != nil {
		if errors.Is(err, ErrResNotFound) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	// add headers
	if h.allowOrigin != nil {
		w.Header().Add("Access-Control-Allow-Origin", *h.allowOrigin)
	}
	w.Header().Add("Content-Type", "application/jrd+json")

	// encode response to json
	ret, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(ret)
}
