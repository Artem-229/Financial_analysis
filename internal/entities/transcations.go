package entities

import "github.com/google/uuid"

type Transaction struct {
	id    uuid.UUID
	item  string
	price int
	class string
}
