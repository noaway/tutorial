package config

import (
	"fmt"
	"io/ioutil"

	"github.com/hashicorp/hcl"
)

type Configuration struct {
	Server     Server           `hcl:"server,block"`
	PostgreSQL PostgreSQLConfig `hcl:"postgre_db,block"`
}

type Server struct {
	HttpAddr string `hcl:"http_addr"`
}

type PostgreSQLConfig struct {
	Host              string `hcl:"host" json:"host"`
	Port              int    `hcl:"port" json:"port"`
	User              string `hcl:"user" json:"user"`
	Password          string `hcl:"password" json:"password"`
	Database          string `hcl:"database" json:"database"`
	SSLMode           string `hcl:"sslmode" json:"sslmode"`
	ShowSQL           bool   `hcl:"showsql" json:"showsql"`
	MaxIdleConnection int    `hcl:"maxidleconnection" json:"maxidleconnection"`
	MaxOpenConnection int    `hcl:"maxopenconnection" json:"maxopenconnection"`
	ApplicationName   string `hcl:"applicationname" json:"applicationname"`
}

func NewConfigure(filename string) *Configuration {
	if filename == "" {
		return nil
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("read config error. [err='%v', filename='%v']", err, filename))
	}
	configure := new(Configuration)
	if err := hcl.Decode(configure, string(data)); err != nil {
		panic(fmt.Sprintf("parse config file error. [err='%v', filename='%v']", err, filename))
	}
	return configure
}
