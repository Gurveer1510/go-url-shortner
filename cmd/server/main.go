package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Gurveer1510/url-shortner/internal/adaptors/persistance"
	"github.com/Gurveer1510/url-shortner/internal/config"
	"github.com/Gurveer1510/url-shortner/internal/interfaces/input/api/rest/handler"
	"github.com/Gurveer1510/url-shortner/internal/interfaces/input/api/rest/routes"
	"github.com/Gurveer1510/url-shortner/internal/usecase"
	"github.com/Gurveer1510/url-shortner/pkg/migrate"
)

func main() {
	database, err := persistance.NewDatabase()
	if err != nil {
		log.Fatalf("failed to connect to database - Error: %s", err)
	}
	log.Println("Connected to Database")

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("error running migrations: %v", err)
	}

	migrate := migrate.NewMigrate(database.GetDB(), cwd+"/migrations")

	err = migrate.RunMigrations()
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			fmt.Println("DB is synced")
		} else {
			log.Fatalf("failed to run migrations - %v", err)
		}
	}

	urlRepo := persistance.NewURLRepo(database)
	urlService := usecase.NewURLService(*urlRepo)
	urlHandler := handler.NewURLHandler(*urlService)

	router := routes.InitRoutes(*urlHandler)

	cfg, err := config.LoadConfig()
	if err != nil {
		panic("port not found")
	}

	err = http.ListenAndServe(fmt.Sprintf(":%v",cfg.APP_PORT), router)
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}

}
