package main

import (
	"SeminarioGoLang/internal/config"
	"SeminarioGoLang/internal/database"
	"SeminarioGoLang/internal/service"
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	configFile := flag.String("config", "./config.yaml", "this is the service config")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	db, err := database.NewDataBase(cfg)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = database.CreateSchema(db)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	srv, err := service.New(db, cfg)
	htppService := service.NewHTTPTransport(srv)

	router := gin.Default()
	htppService.Register(router)
	router.Run()

}
