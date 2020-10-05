package invoices

import "fmt"

func NewErrorUserNotFound(id int) error {
	return fmt.Errorf("User with id #%d not found", id)
}

type User struct {
	Id int `json:"id"`
	Login string `json:"login"`
	Password string `json:"password"`
}
