package http

import (
	"encoding/json"
	"net/http"

	"github.com/adjust/redismq"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/thetinygoat/localeai/pkg/entities/ride"
)

type payload struct {
	ID uuid.UUID `json:"id"`
}

type response struct {
	Data   *payload `json:"data"`
	Status string   `json:"status"`
}

type httpHandler struct {
	queue *redismq.Queue
}

// Routes registers ride http handlers
func Routes(q *redismq.Queue) *chi.Mux {
	r := chi.NewRouter()
	h := httpHandler{queue: q}
	r.Post("/dump", func(rw http.ResponseWriter, r *http.Request) {
		h.dump(rw, r)
	})
	return r
}

func (h *httpHandler) dump(w http.ResponseWriter, r *http.Request) {
	ride := ride.Ride{}
	json.NewDecoder(r.Body).Decode(&ride)
	id := uuid.New()
	ride.JobID = id
	blob, _ := json.Marshal(ride)
	err := h.queue.Put(string(blob))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response{Data: nil, Status: "failed"})
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response{Data: &payload{ID: id}, Status: "success"})
}
