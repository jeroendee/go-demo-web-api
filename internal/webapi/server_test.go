package webapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jeroendee/go-demo-web-api/internal/domain"
	"golang.org/x/exp/slices"
)

// executeRequest, creates a new ResponseRecorder
// then executes the request by calling ServeHTTP in the router
// after which the handler writes the response to the response recorder
// which we can then inspect.
func executeRequest(req *http.Request, s *Server) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)

	return rr
}

func TestRoot(t *testing.T) {
	// Arrange
	s := NewServer()

	req, _ := http.NewRequest("GET", "/", nil)

	// Act
	response := executeRequest(req, s)

	// Assert
	Equal(t, http.StatusOK, response.Code)
	Equal(t, "root.", response.Body.String())
}

func TestPing(t *testing.T) {
	// Arrange
	s := NewServer()

	req, _ := http.NewRequest("GET", "/ping", nil)

	// Act
	response := executeRequest(req, s)

	// Assert
	Equal(t, http.StatusOK, response.Code)
	Equal(t, "pong", response.Body.String())
}

func TestGetClients(t *testing.T) {
	// Arrange
	s := NewServer()

	req, _ := http.NewRequest("GET", "/clients", nil)

	// Act
	response := executeRequest(req, s)

	// Assert
	Equal(t, http.StatusOK, response.Code)
	var c []domain.Client
	err := json.Unmarshal(response.Body.Bytes(), &c)
	if err != nil {
		t.Errorf("unmarshalling failed")
	}
	Equal(t, 100, len(c))
}

func TestGetClient(t *testing.T) {
	// Arrange
	s := NewServer()

	req, _ := http.NewRequest("GET", "/clients/1", nil)

	// Act
	response := executeRequest(req, s)

	// Assert
	Equal(t, http.StatusOK, response.Code)
	var c domain.Client
	err := json.Unmarshal(response.Body.Bytes(), &c)
	if err != nil {
		t.Errorf("unmarshalling failed")
	}
	Equal(t, "1", c.Id)
}

func TestDeleteClient(t *testing.T) {
	// Arrange
	s := NewServer()

	delreq, _ := http.NewRequest("DELETE", "/clients/1", nil)

	// Act
	delresp := executeRequest(delreq, s)

	// Assert
	Equal(t, http.StatusOK, delresp.Code)

	Equal(t, 99, len(s.Db.Clients))
}

func TestUpdateClient(t *testing.T) {
	// Arrange
	s := NewServer()

	c := &domain.Client{
		Id:   "1",
		Name: "test",
	}
	j, err := json.Marshal(c)
	b := bytes.NewReader(j)
	if err != nil {
		t.Errorf("marshalling failed")
	}
	createreq, _ := http.NewRequest("PUT", "/clients/1", b)

	// Act
	createresp := executeRequest(createreq, s)

	// Assert
	Equal(t, http.StatusOK, createresp.Code)

	idx := slices.IndexFunc(s.Db.Clients, func(c *domain.Client) bool { return c.Id == "1" })

	Equal(t, "test", s.Db.Clients[idx].Name)
}

func TestCreateClient(t *testing.T) {
	// Arrange
	s := NewServer()

	c := domain.Client{}
	j, err := json.Marshal(c)
	b := bytes.NewReader(j)
	if err != nil {
		t.Errorf("marshalling failed")
	}
	createreq, _ := http.NewRequest("POST", "/clients", b)

	// Act
	createresp := executeRequest(createreq, s)

	// Assert
	Equal(t, http.StatusCreated, createresp.Code)

	Equal(t, 101, len(s.Db.Clients))
}

func TestCreateClientErr(t *testing.T) {
	// Arrange
	s := NewServer()
	b := io.ReadCloser(http.NoBody)
	createreq, _ := http.NewRequest("POST", "/clients", b)

	// Act
	createresp := executeRequest(createreq, s)

	// Assert
	Equal(t, http.StatusBadRequest, createresp.Code)

	Equal(t, 100, len(s.Db.Clients))
}

func TestClientCtxNotFound(t *testing.T) {
	tcs := []struct {
		urlp string
		want int
	}{
		{"101", http.StatusNotFound},
		{"somestring", http.StatusNotFound},
	}

	for _, tc := range tcs {
		t.Run(fmt.Sprintf("Requesting client with: %s should result in http statuscode: %d", tc.urlp, tc.want), func(t *testing.T) {
			// Arrange
			s := NewServer()

			req, _ := http.NewRequest("GET", "/clients/"+tc.urlp, nil)

			// Act
			response := executeRequest(req, s)

			// Assert
			Equal(t, tc.want, response.Code)
		})
	}
}

func Equal[T comparable](t *testing.T, want, got T) {
	t.Helper()

	if want != got {
		t.Errorf("want: %v; got: %v", want, got)
	}
}
