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

	User *UserPayload `json:"user,omitempty"`

	// We add an additional field to the response here.. such as this
	// elapsed computed property
	Elapsed int64 `json:"elapsed"`
}

func (rd *ClientResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	rd.Elapsed = 10
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

	if resp.User == nil {
		if user, _ := db.GetUser(resp.UserID); user != nil {
			resp.User = NewUserPayloadResponse(user)
		}
	}

	return resp
}

type UserPayload struct {
	*domain.User
	Role string `json:"role"`
}

func NewUserPayloadResponse(user *domain.User) *UserPayload {
	return &UserPayload{User: user}
}
