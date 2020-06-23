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

var conf = config.NewConfigure("/Users/noaway/workspace/tutorial/config.conf")

var Set = wire.NewSet(
	wire.Value(conf),
	provider.ProviderSet,
	wire.Struct(new(service.UserService), "Users"),
	wire.Struct(new(service.ItemService), "*"),
	web.InitWeb,
)

func InitWeb() (http.Handler, error) { panic(wire.Build(Set)) }

type Proc struct {
	httpService svc.HTTPService
}

func (p *Proc) Init() (e error) {
	return nil
}

func (p *Proc) Start() error {
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
