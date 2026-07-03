package entities

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID
	Username string
	Login    string
	Password string
}
