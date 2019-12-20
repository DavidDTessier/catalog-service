package main

import (
	"fmt"
	"os"

	service "github.com/DuffMan2020/catalog-service/service"
	"github.com/cloudfoundry-community/go-cfenv"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	appEnv, err := cfenv.Current()

	if err != nil {
		fmt.Println("CF Environment not detected.")
		server := service.NewServer()
		server.Run(":" + port)

	} else {
		server := service.NewServerFromCFEnv(appEnv)
		server.Run(":" + port)

	}

}
