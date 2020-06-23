package provider

import (
	"github.com/noaway/tutorial/service"
	"github.com/noaway/tutorial/config"
)

type Item struct{
	service.Users
	Conf *config.Configuration
}

func(i * Item) ShowName() string{
	return i.GetUser().Name + " "+ i.Conf.Server.HttpAddr
}