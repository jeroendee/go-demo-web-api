package domain

import (
	"strconv"

	"github.com/brianvoe/gofakeit/v6"
)

func GetClients() []*Client {
	c := createClientsData()
	return c
}

// iterate 100 times over the creation of client data and add that to the slice
func createClientsData() []*Client {
	var l []*Client

	for i := 0; i < 100; i++ {
		c := &Client{
			Id:        strconv.Itoa(i + 1),
			Name:      gofakeit.FirstName(),
			Surname:   gofakeit.LastName(),
			Birthdate: gofakeit.Date().Format("2006-01-02"),
			Email:     gofakeit.Email(),
			Phone:     gofakeit.Phone(),
			Age:       int32(gofakeit.Number(10, 100)),
		}
		l = append(l, c)
	}

	return l
}
