package main

import (
	config "interview/pkg/configs"
	"interview/pkg/core/common"
	"interview/pkg/core/service"
	"interview/pkg/infrastructure/db"
	"interview/pkg/infrastructure/repository"
	"interview/pkg/interface/api"
	"interview/pkg/logger"
)

func main() {
	logger.Init("interview")
	lg := logger.General.Component("main")
	config.Load()
	cfg := config.Get()

	dbRepo := repository.NewRepository(db.GetDatabase(false, cfg.DBConnectionString()))

	cartService := service.NewCartService(service.CartServiceParam{
		CartItemRepo:   dbRepo,
		CartEntityRepo: dbRepo,
	})

	apiServer := api.NewServer(&api.ServerConfig{
		Port:    cfg.Server.Port,
		GinMode: cfg.Server.GinMode,
		Services: common.Services{
			Cart: cartService,
		},
	})
	lg.Println("server started")
	// start the API server
	if err := apiServer.Start(); err != nil {
		lg.WithError(err).Fatal("failed to start API server")
	}

}
