package main

import (
	"bigtable_api/database"
	"bigtable_api/handlers"
	"bigtable_api/repository"
	"bigtable_api/router"
	"bigtable_api/server"
	"bigtable_api/usecase"
	"context"
	"log"
)

const port = "7000"

func main() {
	ctx := context.Background()
	clientInstance, err := database.GetClientSingleton(ctx)
	if err != nil {
		log.Fatalln("error creating database instance. Error: ", err.Error())
	}

	climateRepo := repository.NewClimateRepository(clientInstance)

	climateUsecase := usecase.NewClimateUsecase(climateRepo)

	climateHandler := handlers.NewClimateHandler(climateUsecase)

	router := router.InitializeRouter(climateHandler)

	server := server.NewServer(":"+port, router)
	server.Start()
}
