package main

import (
	"flag"
	"log"

	"github.com/daheige/inject-demo/internal/application/service"
	"github.com/daheige/inject-demo/internal/config"
	"github.com/daheige/inject-demo/internal/infras/inject"
	"github.com/daheige/inject-demo/internal/infras/persistence"
	"github.com/daheige/inject-demo/internal/interfaces"
)

var cfgFile string

func init() {
	flag.StringVar(&cfgFile, "config_file", "./app.yaml", "config file")
	flag.Parse()
}

func main() {
	config.CfgFile = cfgFile
	di := inject.New()
	di.BuildProvides(
		config.LoadConfig,
		config.InitAppConf,
		config.InitDBConf,
		persistence.NewUserRepository,
		service.NewUserService,
		interfaces.NewServer,
	)

	err := di.Invoke(func(server *interfaces.Server) {
		server.Run()
	})

	if err != nil {
		log.Fatalln("server run error: ", err)
	}
}
