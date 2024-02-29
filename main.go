package main

import (
	"fmt"
	"intikom_test/config"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viperConfig := config.InitViper()
	dbInit := config.InitDBConnection(viperConfig)
	ginInit := gin.Default()

	config.Bootstrap(&config.BootstrapConfig{
		DB:     dbInit,
		App:    ginInit,
		Config: viperConfig,
	})

	webPort := viper.GetString("web.port")

	ginInit.Run(fmt.Sprintf(":%s", webPort))
}
