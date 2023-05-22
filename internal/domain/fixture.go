package domain

func GetClients() []*Client {
	c := []*Client{
		{Id: "1", UserID: 100, Name: "John", Surname: "Connor", Birthdate: "01-01-1980"},
		{Id: "2", UserID: 200, Name: "Sarah", Surname: "Connor", Birthdate: "01-01-1960"},
		{Id: "3", UserID: 300, Name: "Arnold", Surname: "Schwarzenegger", Birthdate: "01-01-1950"},
	}
	return c
}

func GetUsers() []*User {
	u := []*User{
		{ID: 100, Name: "John"},
		{ID: 200, Name: "Sarah"},
		{ID: 300, Name: "Sarah"},
	}
	return u
}
