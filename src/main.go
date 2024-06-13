package main

import (
	"github.com/nazarow/golang-clean/api"
	"github.com/nazarow/golang-clean/config"
	"github.com/nazarow/golang-clean/data/cache"
	"github.com/nazarow/golang-clean/data/db"
	"github.com/nazarow/golang-clean/data/db/migrations"
	"github.com/nazarow/golang-clean/pkg/logging"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {

	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)
	// logging.NewLogger(cfg)

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}

	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	migrations.Up_1()

	api.InitServer(cfg)
}
