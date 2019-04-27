package main

import (
	"github.com/agustin-sarasua/user-system/src/config"
	"github.com/agustin-sarasua/user-system/src/logger"
	"github.com/agustin-sarasua/user-system/src/routers"
	"github.com/agustin-sarasua/user-system/src/service"
)

func init() {
	config.Cfg = config.NewConfiguration("DEVELOPMENT")
	service.CreateServices()
}

func main() {
	logger.Info("Starting User System API...", nil)
	router = routers.CreateRouter()
	router.Run()
}
