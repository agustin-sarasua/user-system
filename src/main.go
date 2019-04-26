package main

import (
	"github.com/agustin-sarasua/user-system/src/logger"
	"github.com/agustin-sarasua/user-system/src/routers"
)

func init() {

}

func main() {
	logger.Info("Starting User System API...", nil)
	router = routers.CreateRouter()
	router.Run()
}
