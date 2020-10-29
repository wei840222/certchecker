package main

import (
	"github.com/wei840222/certchecker/bot"
	"github.com/wei840222/certchecker/cert"
)

func main() {
	go cert.StartCertCheck()
	bot.HandleUpdate()
}
