package server

import (
	"github.com/containers/libpod/pkg/api/handlers"
	"github.com/gorilla/mux"
)

func (s *APIServer) registerDistributionHandlers(r *mux.Router) error {
	r.HandleFunc(VersionedPath("/distribution/{name}/json"), handlers.UnsupportedHandler)
	// Added non version path to URI to support docker non versioned paths
	r.HandleFunc("/distribution/{name}/json", handlers.UnsupportedHandler)
	return nil
}
