package brands

import (
	"errors"
)

type UserType = string

const (
	Admin   UserType = "ADMIN"
	General UserType = "USER"
)

type (
	GeneralUser struct {
		ID       string   `json:"id"`
		Name     string   `json:"name"`
		UserType UserType `json:"user_type"`
	}

	AdminUser GeneralUser

	User interface {
		GeneralUser | AdminUser
	}
)

func IsAdmin(u GeneralUser) (*AdminUser, error) {
	if u.UserType == Admin {
		adminUser := AdminUser(u)
		return &adminUser, nil
	}

	return nil, errors.New("User is not an admin")
}
