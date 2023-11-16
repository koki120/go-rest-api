package entity

import (
	"time"

	"github.com/koki120/go-rest-api/domain/entconst"
)

type User struct {
	UserID         string
	Name           string
	HashedPassword string
	UserType       entconst.UserType
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}

type UserCrate struct {
	UserID   string
	Name     string
	Password string
	UserType entconst.UserType
}
