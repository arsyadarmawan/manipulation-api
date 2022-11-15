package main

import (
	"fmt"
	"net/http"
	"os"
	"task/app"
	"task/controller"
	"task/exception"
	"task/helper"
	"task/repository"
	"task/service"

	_ "github.com/lib/pq"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	validate := validator.New()
	error := godotenv.Load(".env")
	helper.PanicHandling(error)

	configDB := app.ConfigDB{
		DB_Username: os.Getenv("DB_USERNAME"),
		DB_Password: os.Getenv("DB_PASSWORD"),
		DB_Host:     os.Getenv("DB_HOST"),
		DB_Port:     os.Getenv("DB_PORT"),
		DB_Database: os.Getenv("DB_NAME"),
	}

	db := app.NewDB(&configDB)

	// tasks

	// users
	repoUser := repository.NewAuthRepo(db)
	serviceUser := service.NewUserService(repoUser, validate)
	controllerUser := controller.NewAuthController(serviceUser)

	router := app.NewRouter(&controllerUser)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    os.Getenv("APP_URL"),
		Handler: router,
	}
	fmt.Printf("Running on server " + os.Getenv("APP_URL") + "\n")
	err := server.ListenAndServe()
	helper.PanicHandling(err)

}
