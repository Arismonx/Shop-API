package databases

import (
	"fmt"
	"log"
	"sync"

	"github.com/Arismonx/shop-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDatabase struct {
	*gorm.DB
}

var (
	postgreaDatabaseInstance *postgresDatabase
	once                     sync.Once
)

func NewPostgresDatabase(config *config.Database) Database {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s search_path=%s",
			config.Host,
			config.Port,
			config.User,
			config.Password,
			config.DBname,
			config.SSLmode,
			config.Schema,
		)
		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			panic(err)
		}

		log.Printf("Connected to database %s", config.DBname)

		postgreaDatabaseInstance = &postgresDatabase{conn}
	})
	return postgreaDatabaseInstance
}

func (db *postgresDatabase) ConnectionGetting() *gorm.DB {
	return db.DB
}
