package main

import (
	"github.com/wei840222/certchecker/bot"
	"github.com/wei840222/certchecker/cert"
	"github.com/wei840222/certchecker/handler"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {

	go cert.StartCertCheck()
	go bot.StartAlert()
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
