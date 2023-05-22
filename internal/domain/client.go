package domain

type Client struct {
	Id        string `json:"id"`
	UserID    int64  `json:"user_id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Birthdate string `json:"birthdate"`
}
