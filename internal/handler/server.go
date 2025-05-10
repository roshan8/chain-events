package handler

import (
	"net/http"

	"github.com/roshan8/change-events/api"
	"github.com/roshan8/change-events/internal/service"
)

var _ api.ServerInterface

type Server struct {
	// Add any dependencies here
	kubernetesService *service.KubernetesService
}

// NewServer returns a new Server
func NewServer(kubernetesService *service.KubernetesService) *Server {
	return &Server{
		kubernetesService: kubernetesService,
	}
}

// Implement the generated ServerInterface
func (s *Server) CreateKubernetesEvent(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()

	// var kevent api.KubernetesEvent
	// if err := json.NewDecoder(r.Body).Decode(&kevent); err != nil {
	// 	writeError(w, http.StatusBadRequest, "Invalid request body")
	// 	return
	// }

	// createdKEvent, err := s.kubernetesService.CreateKubernetesEvent(ctx, kevent)
	// if err != nil {
	// 	writeError(w, http.StatusInternalServerError, "Failed to create pet")
	// 	return
	// }

	// // Write the response
	// api.WriteJSON(w, http.StatusCreated, createdKEvent)
}

func (s *Server) ListKubernetesEvents(w http.ResponseWriter, r *http.Request, params api.ListKubernetesEventsParams) {

}

func (s *Server) CreateGCPEvent(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) ListGCPEvents(w http.ResponseWriter, r *http.Request, params api.ListGCPEventsParams) {

}
