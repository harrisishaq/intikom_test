package main

import (
	"fmt"
	"intikom_test/config"
	"log"

	"github.com/gin-gonic/gin"
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

	webPort := viperConfig.GetString("web.port")
	log.Println(webPort)

	ginInit.Run(fmt.Sprintf(":%s", webPort))
}
