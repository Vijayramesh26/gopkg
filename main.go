package main

import (
	"fmt"
	"gopkg/apps/common"
	"gopkg/config"
	"gopkg/db"
	"gopkg/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the logger for app-wide logging
	utils.InitLogger()
	logger := &utils.Logger{}
	logger.SetReqID()

	// db.Init(logger)
	db.GlobalDBInit(logger)
	config.Init(logger)

	// Set up the router
	router := mux.NewRouter()

	// Define the /hello route (GET method)
	router.HandleFunc("/hello", common.Ready).Methods(http.MethodGet)

	// Start the server
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
