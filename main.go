package main

import (
	"github.com/wei840222/certchecker/bot"
	"github.com/wei840222/certchecker/cert"
	"github.com/wei840222/certchecker/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	go cert.StartCertCheck()
	go bot.HandleUpdate()

	r := gin.Default()
	d := r.Group("/domain")
	{
		d.GET("", handler.ListDomain)
		d.POST("", handler.CreateDomain)
	}
	r.Run()
}
