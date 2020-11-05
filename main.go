package main

import (
	"fmt"

	"github.com/wei840222/certchecker/bot"
	"github.com/wei840222/certchecker/cert"
	"github.com/wei840222/certchecker/handler"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./conf") // optionally look for config in the working directory
	AutomaticEnv()
	SetEnvKeyReplacer(".", "_")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	go cert.StartCertCheck()
	go bot.HandleUpdate()

	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./web/dist", false)))
	d := r.Group("/domain")
	{
		d.GET("", handler.ListDomain)
		d.POST("", handler.CreateDomain)
		d.DELETE(":id", handler.DeleteDomain)
	}
	r.Run()
}
