package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/network-api/api"
	"github.com/network-api/config"
	"github.com/network-api/fabric"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	sdk, err := fabric.InitializeSDK(config.FabricSDKConfig)
	if err != nil {
		log.Fatal("Failed to initialize Fabric SDK:", err)
	}
	defer sdk.Close()

	router := mux.NewRouter()
	api.RegisterRoutes(router, sdk)

	log.Println("Server is running on port", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, router))
}