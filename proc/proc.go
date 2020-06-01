//+build wireinject

package proc

import (
	"net"
	"net/http"

	"github.com/google/wire"
	"github.com/noaway/tutorial/config"
	"github.com/noaway/tutorial/internal/svc"
	"github.com/noaway/tutorial/provider"
	"github.com/noaway/tutorial/service"
	"github.com/noaway/tutorial/web"
	"github.com/sirupsen/logrus"
)

var Set = wire.NewSet(
	config.NewConfigure,
	wire.Struct(new(provider.UserProvider)),
	wire.Bind(new(service.Users), new(*provider.UserProvider)),
	wire.Struct(new(service.UserService), "Users"),
	web.InitWeb,
)

func InitWeb() (http.Handler, error) { panic(wire.Build(Set)) }

func GetConfig(filename string) *config.Configuration { panic(wire.Build(Set)) }

type Proc struct {
	httpService svc.HTTPService
}

func (p *Proc) Init() (e error) {
	return nil
}

func (p *Proc) Start() error {
	conf := GetConfig("/Users/noaway/workspace/tutorial/config.conf")
	httpListener, err := net.Listen("tcp", conf.Server.HttpAddr)
	if err != nil {
		return err
	}

	logrus.Info("addr: ", conf.Server.HttpAddr)
	handler, err := InitWeb()
	if err != nil {
		return err
	}

	return p.httpService.Service(httpListener, handler)
}

func (p *Proc) Stop() error {
	p.httpService.Shutdown()
	return nil
}
