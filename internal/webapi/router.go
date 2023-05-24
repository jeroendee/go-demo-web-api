package webapi

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/jeroendee/go-demo-web-api/internal/domain"
)

func (s *Server) SetupRouter() {
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(middleware.URLFormat)
	s.Router.Use(render.SetContentType(render.ContentTypeJSON))

	s.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})

	s.Router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	// RESTy routes for "clients" resource
	s.Router.Route("/clients", func(r chi.Router) {
		r.Get("/", s.ListClients)
		r.Post("/", s.CreateClient)
		r.Route("/{clientId}", func(r chi.Router) {
			r.Use(s.ClientCtx)            // Load the *Client on the request context
			r.Get("/", s.GetClient)       // GET /clients/123
			r.Put("/", s.UpdateClient)    // PUT /clients/123
			r.Delete("/", s.DeleteClient) // DELETE /clients/123
		})
	})
}

func (s *Server) ListClients(w http.ResponseWriter, r *http.Request) {
	if err := render.RenderList(w, r, NewClientListResponse(s.Db.Clients)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

// GetClient returns the specific Client. You'll notice it just
// fetches the Client right off the context, as its understood that
// if we made it this far, the Client must be on the context. In case
// its not due to a bug, then it will panic, and our Recoverer will save us.
func (s *Server) GetClient(w http.ResponseWriter, r *http.Request) {
	// Assume if we've reach this far, we can access the client
	// context because this handler is a child of the ClientCtx
	// middleware. The worst case, the recoverer middleware will save us.
	client := r.Context().Value("client").(*domain.Client)

	if err := render.Render(w, r, NewClientResponse(client)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

// CreateClient persists the posted Client and returns it
// back to the client as an acknowledgement.
func (s *Server) CreateClient(w http.ResponseWriter, r *http.Request) {
	data := &ClientRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	client := data.Client
	s.Db.NewClient(client)

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewClientResponse(client))
}

// UpdateClient updates an existing Client in our persistent store.
func (s *Server) UpdateClient(w http.ResponseWriter, r *http.Request) {
	client := r.Context().Value("client").(*domain.Client)

	data := &ClientRequest{Client: client}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	client = data.Client
	s.Db.UpdateClient(client.Id, client)

	render.Render(w, r, NewClientResponse(client))
}

// DeleteClient removes an existing Client from our persistent store.
func (s *Server) DeleteClient(w http.ResponseWriter, r *http.Request) {
	var err error

	// Assume if we've reach this far, we can access the Client
	// context because this handler is a child of the ClientCtx
	// middleware. The worst case, the recoverer middleware will save us.
	Client := r.Context().Value("client").(*domain.Client)

	Client, err = s.Db.RemoveClient(Client.Id)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	render.Render(w, r, NewClientResponse(Client))
}

// ClientCtx middleware is used to load an Client object from
// the URL parameters passed through as the request. In case
// the Client could not be found, we stop here and return a 404.
func (s *Server) ClientCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var client *domain.Client
		var err error

		c := chi.URLParam(r, "clientId")

		// This will be handled by /clients call
		if c == "" {
			render.Render(w, r, ErrNotFound)
			return
		}

		client, err = s.Db.GetClient(c)

		if err != nil {
			render.Render(w, r, ErrNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "client", client)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
