package main

import (
	"github.com/fxivan/routeOptimizer/config"
	"github.com/fxivan/routeOptimizer/database"
	"github.com/fxivan/routeOptimizer/server"
)

func main() {
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)
	server.NewEchoServer(conf, db).Start()
}
