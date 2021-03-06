package mysql

import (
	"github.com/flannel-dev-lab/RBAC/vars"
	_ "github.com/go-sql-driver/mysql" // Importing mysql Driver
)

// AddOperation Adds a new operation
func (databaseService *DatabaseService) AddOperation(name, description string) (operation vars.Operation, err error) {
	stmt, err := databaseService.Conn.Prepare("INSERT INTO `rbac_operation` SET `name`= ?, description = ?")
	if err != nil {
		return operation, err
	}

	result, err := stmt.Exec(name, description)
	if err != nil {
		return operation, err
	}

	insertId, insertIdErr := result.LastInsertId()
	if insertIdErr != nil {
		return operation, insertIdErr
	}

	operation.Id = int(insertId)
	operation.Name = name
	operation.Description = description

	return operation, nil
}

// DeleteOperation Deletes an existing operation
func (databaseService *DatabaseService) DeleteOperation(operationName string) (bool, error) {
	stmt, prepErr := databaseService.Conn.Prepare("DELETE FROM `rbac_operation` WHERE `name` = ?")
	if prepErr != nil {
		return false, prepErr
	}

	_, err := stmt.Exec(operationName)
	if err != nil {
		return false, err
	}

	return true, nil
}
