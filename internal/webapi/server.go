package webapi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jeroendee/go-demo-web-api/internal/datastore"
	"github.com/jeroendee/go-demo-web-api/internal/domain"
)

// Global available data
var db = &datastore.Db{
	Clients: domain.GetClients(),
	Users:   domain.GetUsers(),
}

// What does the Server need in order to do it's job
// No global state
type Server struct {
	Router *chi.Mux
	// Db     *datastore.Data ? server contained data?
}

func NewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	s.SetupRouter()
	return s
}

// Make `server` a `http.Handler`
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}
