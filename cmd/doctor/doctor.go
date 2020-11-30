package main

import (

	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-delve/delve/pkg/config"
	"github.com/jmoiron/sqlx"
	
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

	service, err := doctor.New(db, cfg)
	htppService := service.NewHTTPTransport(service)
	
	router := gin.Default()
	htppService.Register(router)
	router.Run()

}