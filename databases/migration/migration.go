package main

import (
	"fmt"

	"github.com/Arismonx/shop-api/config"
	"github.com/Arismonx/shop-api/databases"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)

	fmt.Println(db.ConnectionGetting())
}
