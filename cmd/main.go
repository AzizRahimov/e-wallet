package main

import (
	"github.com/AzizRahimov/e-wallet/db"
	routes "github.com/AzizRahimov/e-wallet/pkg/routes"
	"github.com/AzizRahimov/e-wallet/utils"
)

func main() {
	utils.ReadSettings()
	dbConnection := db.StartDbConnection()
	routes.RunServerAndRoutes(dbConnection)
}
