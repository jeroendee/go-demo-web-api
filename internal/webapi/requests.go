package webapi

import (
	"errors"
	"net/http"
	"strings"

	"github.com/jeroendee/go-demo-web-api/internal/domain"
)

type ClientRequest struct {
	*domain.Client
}

func (c *ClientRequest) Bind(r *http.Request) error {
	// c.Client is nil if no Client fields are sent in the request. Return an
	// error to avoid a nil pointer dereference.

	if c.Client == nil {
		return errors.New("missing required Client fields")
	}

	// just a post-process after a decode..
	c.Client.Name = strings.ToLower(c.Client.Name) // as an example, we down-case
	return nil
}
