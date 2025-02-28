package main

import (
	"github.com/Arismonx/shop-api/config"
	"github.com/Arismonx/shop-api/databases"
	"github.com/Arismonx/shop-api/server"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db.ConnectionGetting())
	server.Start()
}
