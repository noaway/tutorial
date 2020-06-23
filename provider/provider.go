package provider

import (
	"github.com/google/wire"
	"github.com/noaway/tutorial/service"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(UserProvider)),
	wire.Bind(new(service.Users), new(*UserProvider)),
	wire.Struct(new(Item), "*"),
	wire.Bind(new(service.ItemServices), new(*Item)),
)

type UserProvider struct{}

func (up *UserProvider) GetUser() *service.User {
	return &service.User{ID: 1, Name: "noaway", Age: 22}
}
