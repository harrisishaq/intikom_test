package config

import (
	"intikom_test/entity"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB     *gorm.DB
	App    *gin.Engine
	Config *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	// setup services
	// setup controller
	// migrate database
	migrate := gormigrate.New(config.DB, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "INIT_MIGRATION",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(
					&entity.User{},
					&entity.Task{},
				)
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&entity.User{}, &entity.Task{})
			},
		},
	})

	if err := migrate.Migrate(); err != nil {
		log.Fatalf("Migrate failed: %+v\n", err)
	}
}
