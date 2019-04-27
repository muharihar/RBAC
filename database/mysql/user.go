package mysql

import (
	"github.com/flannel-dev-lab/RBAC/vars"
)

// (RC-04) Core RBAC: Creates a new RBAC user.  The User will not carry any sessions during the creation
func (databaseService *DatabaseService) AddUser(name string) (user vars.User, err error) {

	stmt, err := databaseService.Conn.Prepare("INSERT INTO `rbac_user` SET `name`= ?")
	if err != nil {
		return user, err
	}

	result, err := stmt.Exec(name)
	if err != nil {
		return user, err
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return user, err
	}

	user.Id = int(insertId)
	user.Name = name

	return user, nil
}

// (RC-26) Core RBAC: Deletes an existing user from RBAC, Deletes Sessions and User assignments
func (databaseService *DatabaseService) DeleteUser(userId int) (bool, error) {
	// TODO Delete User Assignments and Delete Sessions
	stmt, err := databaseService.Conn.Prepare("DELETE FROM `rbac_user` WHERE `rbac_user_id`= ?")

	if err != nil {
		return false, err
	}

	_, err = stmt.Exec(userId)
	if err != nil {
		return false, err
	}

	return true, nil
}

// (RC-09) Core RBAC: Returns a set of roles assigned to a given user
func (databaseService *DatabaseService) AssignedRoles(userId int) ([]vars.Role, error) {
	stmt, prepErr := databaseService.Conn.Prepare("SELECT `rbac_role_id` FROM `rbac_user_role` WHERE `rbac_user_id` = ?")
	if prepErr != nil {
		return nil, prepErr
	}

	result, err := stmt.Query(userId)
	if err != nil {
		return nil, err
	}

	var roles []vars.Role
	for result.Next() {
		var role vars.Role
		err = result.Scan(&role.Id)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}


