package main

import (
	"fmt"
	"os"
	"os/signal"

	server "github.com/DarkSoul94/film-rest/cmd/httpserver"
	"github.com/DarkSoul94/film-rest/config"
	"github.com/DarkSoul94/film-rest/pkg/logger"
)

func main() {
	conf := config.InitConfig()
	logger.InitLogger(conf)

	apphttp := server.NewApp(conf)
	go apphttp.Run(conf)

	fmt.Println(
		fmt.Sprintf(
			"Service %s is running",
			conf.AppName,
		),
	)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	apphttp.Stop()

	fmt.Println(
		fmt.Sprintf(
			"Service %s is stopped",
			conf.AppName,
		),
	)
}
