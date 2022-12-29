package main

import (
	"github.com/AzizRahimov/e-wallet/db"
	routes "github.com/AzizRahimov/e-wallet/pkg/routes"
	"github.com/AzizRahimov/e-wallet/utils"
)

func main() {
	utils.ReadSettings()                    // read config from file
	dbConnection := db.StartDbConnection()  // connect to db
	routes.RunServerAndRoutes(dbConnection) // run server and routes
}
