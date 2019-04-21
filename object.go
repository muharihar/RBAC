package RBAC

import ( 
    "errors"
)

// An object can be any system resource subject to access control
type Object struct {
    Id              int
    Name            string
    Description     string
}


// (RC-33) Core RBAC: Returns the set of operations a given role
// is permitted to perform on a given object
func RoleOperationsOnObject(role Role, object Object) ([]int, error) {
    return nil, errors.New("Not yet implemented")
}

// (RC-42) Core RBAC: Returns the set of operations a given user
// is permitted to perform on a give object
func UserOperationsOnObject(user User, object Object) ([]int, error) {
    return nil, errors.New("Not yet implemented")
}

// Create an Object
func CreateObject(name string, description string) (Object, error) {
    var object Object

    DbInit()
    
    stmt, prepErr := DBWrite.Prepare("INSERT INTO `rbac_object` SET `name`= ?, description = ?")
    if prepErr != nil {
        return object, prepErr
    }

    result, err := stmt.Exec(name, description)
    if err != nil {
        return object, err
    }

    insertId, insertIdErr := result.LastInsertId()
    if insertIdErr != nil {
        return object, insertIdErr
    }

    object.Id = int(insertId)
    object.Name = name
    object.Description = description

    return object, nil
}

// Remove an Object
func RemoveObject(object Object) (bool, error) {
    DbInit()

    stmt, prepErr := DBWrite.Prepare("DELETE FROM `rbac_object` WHERE `rbac_object_id` = ?")
    if prepErr != nil {
        return false, prepErr
    }

    _, err := stmt.Exec(object.Id)
    if err != nil {
        return false, err
    }

    return true, nil
}