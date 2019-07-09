package domain

type UserData struct {
	Name string  `json:"name"`
	Email string `json:"email"`
}

type User struct {
	Id string `json:"id"`
	Name string  `json:"name"`
	Email string `json:"email"`
}
