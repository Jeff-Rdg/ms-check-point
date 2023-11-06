package main

import (
	"ms-control-point/configs"
	"ms-control-point/internal/handlers"
	"ms-control-point/internal/infra/database"
	"net/http"
)

func main() {

	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := config.LoadDatabase()
	if err != nil {
		panic(err)
	}

	userDb := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDb)

	r := config.LoadRoutes(userHandler)

	http.ListenAndServe(":8000", r)
}
