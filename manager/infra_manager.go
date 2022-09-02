package manager

import (
	"log"
	"os"

	"github.com/wisnuragaprawida/golang-rest-api/config"
	"github.com/wisnuragaprawida/golang-rest-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Infra disini bertugas sebagai database penyimpanan pengganti slice
type Infra interface {
	SqlDb() *gorm.DB
}

type infra struct {
	db *gorm.DB
}

func (i *infra) SqlDb() *gorm.DB {
	return i.db
}

func NewInfra(config config.Config) Infra {
	resource, err := initDbResource(config.DataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &infra{db: resource}
}

func initDbResource(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})

	env := os.Getenv("ENV")
	if env == "migration" {
		db.Debug()
		db.AutoMigrate(&model.Product{})
		db.AutoMigrate(&model.User{})
	} else if env == "dev" {
		db.Debug()
	}
	if err != nil {
		return nil, err
	}
	return db, nil
}
