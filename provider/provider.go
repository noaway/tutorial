package provider

import (
	"github.com/noaway/tutorial/service"
)

type UserProvider struct{}

func (up *UserProvider) GetUser() *service.User {
	return &service.User{ID: 1, Name: "noaway", Age: 22}
}
