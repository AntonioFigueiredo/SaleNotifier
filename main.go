package main

import (
	"fmt"

	"github.com/davidclarafigueiredo/SaleNotifier/config"
	"github.com/davidclarafigueiredo/SaleNotifier/connect"
	"github.com/davidclarafigueiredo/SaleNotifier/handler"
)

func main() {
	config.Init()
	fmt.Printf("%s\n", handler.GetPrice(connect.Connect()))
	//handler.SendMail()
}
