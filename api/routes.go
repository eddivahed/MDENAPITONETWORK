// /api/routes.go
package api

import (
	"github.com/gorilla/mux"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/network-api/api/handlers"
)

func RegisterRoutes(router *mux.Router, sdk *fabsdk.FabricSDK) {
	router.Use(LoggingMiddleware)
	router.HandleFunc("/register", handlers.RegisterUser()).Methods("POST")
	router.HandleFunc("/invoke", handlers.InvokeChaincode()).Methods("POST")
	router.HandleFunc("/query", handlers.QueryChaincode()).Methods("POST")
}