package webapi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jeroendee/go-demo-web-api/internal/datastore"
	"github.com/jeroendee/go-demo-web-api/internal/domain"
)

// The server needs a ...
type Server struct {
	Router *chi.Mux
	Db     *datastore.Db
}

func NewServer() *Server {
	s := &Server{}
	s.Db = &datastore.Db{
		Clients: domain.GetClients(),
	}
	s.Router = chi.NewRouter()
	s.SetupRouter()
	return s
}

// Make `server` a `http.Handler`
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}
