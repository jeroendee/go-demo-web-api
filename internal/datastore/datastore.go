package datastore

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/jeroendee/go-demo-web-api/internal/domain"
)

type Db struct {
	Clients []*domain.Client
	Users   []*domain.User
}

func (db *Db) NewClient(client *domain.Client) (string, error) {
	client.Id = fmt.Sprintf("%d", rand.Intn(100)+10)
	db.Clients = append(db.Clients, client)
	return client.Id, nil
}

func (db *Db) GetClient(id string) (*domain.Client, error) {
	for _, c := range db.Clients {
		if c.Id == id {
			return c, nil
		}
	}
	return nil, errors.New("client not found")
}

func (db *Db) UpdateClient(id string, client *domain.Client) (*domain.Client, error) {
	for i, c := range db.Clients {
		if c.Id == id {
			db.Clients[i] = client
			return c, nil
		}
	}
	return nil, errors.New("client not found")
}

func (db *Db) RemoveClient(id string) (*domain.Client, error) {
	for i, c := range db.Clients {
		if c.Id == id {
			db.Clients = append((db.Clients)[:i], (db.Clients)[i+1:]...)
			return c, nil
		}
	}
	return nil, errors.New("client not found")
}

func (db *Db) GetUser(id int64) (*domain.User, error) {
	for _, u := range db.Users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, errors.New("user not found")
}
