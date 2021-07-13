package main

import (
	"log"

	"github.com/daheige/inject-demo/internal/application/service"
	"github.com/daheige/inject-demo/internal/config"
	"github.com/daheige/inject-demo/internal/infras/persistence"
	"github.com/daheige/inject-demo/internal/interfaces"
)

func main() {
	// read config
	config.CfgFile = "./app.yaml"
	viperEntry, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("read config error:%s\n", err.Error())
	}

	appConf, err := config.InitAppConf(viperEntry) // init app config
	if err != nil {
		log.Fatalf("init app config error:%s\n", err.Error())
	}

	db, err := config.InitDBConf(viperEntry) // init db instance
	if err != nil {
		log.Fatalf("init db instance error:%s\n", err.Error())
	}

	// create user repo
	userRepo := persistence.NewUserRepository(db)

	// create user service
	userService := service.NewUserService(appConf, userRepo)

	// create http server
	server := interfaces.NewServer(appConf, userService)

	// run server
	server.Run()
}
