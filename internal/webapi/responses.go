package webapi

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/jeroendee/go-demo-web-api/internal/domain"
)

// ClientResponse is the response payload for the Client data model.
// See NOTE above in ClientRequest as well.
//
// In the ClientResponse object, first a Render() is called on itself,
// then the next field, and so on, all the way down the tree.
// Render is called in top-down order, like a http handler middleware chain.
type ClientResponse struct {
	*domain.Client
}

func (rd *ClientResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

func NewClientListResponse(clients []*domain.Client) []render.Renderer {
	list := []render.Renderer{}
	for _, client := range clients {
		list = append(list, NewClientResponse(client))
	}
	return list
}

func NewClientResponse(client *domain.Client) *ClientResponse {
	resp := &ClientResponse{Client: client}

	return resp
}
