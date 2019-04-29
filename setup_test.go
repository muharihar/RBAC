package RBAC

import (
	"github.com/flannel-dev-lab/RBAC/database"
	"log"
	"os"
)

// Role test parameters
var roleObject RoleObject

func setupRoleTest() {
	dbService, err := database.CreateDatabaseObject("mysql")
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = dbService.CreateDBConnection(
		os.Getenv("RBAC_DB_DRIVER"),
		os.Getenv("RBAC_DB_USERNAME"),
		os.Getenv("RBAC_DB_PASSWORD"),
		os.Getenv("RBAC_DB_HOSTNAME"),
		os.Getenv("RBAC_DB_NAME"),
		os.Getenv("RBAC_DB_PORT"))

	if err != nil {
		log.Fatalf(err.Error())
	}

	roleObject.DBService = dbService
}

func tearDownRoleTest() {
	err := roleObject.DBService.CloseConnection()
	if err != nil {
		log.Fatalf(err.Error())
	}
}

// User test parameters

// Session test parameters
var sessionObject SessionObject

func setupSessionTest() {
	dbService, err := database.CreateDatabaseObject("mysql")
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = dbService.CreateDBConnection(
		os.Getenv("RBAC_DB_DRIVER"),
		os.Getenv("RBAC_DB_USERNAME"),
		os.Getenv("RBAC_DB_PASSWORD"),
		os.Getenv("RBAC_DB_HOSTNAME"),
		os.Getenv("RBAC_DB_NAME"),
		os.Getenv("RBAC_DB_PORT"))

	if err != nil {
		log.Fatalf(err.Error())
	}

	sessionObject.DBService = dbService
}

func tearDownSessionTest() {
	err := sessionObject.DBService.CloseConnection()
	if err != nil {
		log.Fatalf(err.Error())
	}
}

