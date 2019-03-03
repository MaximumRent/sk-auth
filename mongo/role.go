package mongo

import (
	"context"
	"errors"
	"gopkg.in/mgo.v2/bson"
	"log"
	"sk-auth/auth/entity"
	"strings"
)

func hasAccess(requestedAccess, roleAccess string) bool {
	if roleAccess == _FULL_ACCESS_PATH {
		return true
	}
	requestedAccessPieces := strings.Split(requestedAccess, _ROLE_SEPARATOR)
	roleAccessPieces := strings.Split(roleAccess, _ROLE_SEPARATOR)
	if len(requestedAccessPieces) != 2 || len(roleAccessPieces) != 2 {
		log.Printf("Invalid requested access - %s ; or role access - %s")
		return false
	}
	if requestedAccessPieces[0] != roleAccessPieces[0] {
		return false
	}
	if (requestedAccessPieces[1] == roleAccessPieces[1]) || (roleAccessPieces[1] == "*") {
		return true
	}
	return false
}

func AddPathToRole(roleName string, path string) error {
	collection := client.Database(SK_DB_NAME).Collection(ROLES_COLLECTION_NAME)
	filter := bson.M {
		"name": roleName,
	}
	update := bson.M {
		"$push": bson.M {
			"path": path,
		},
	}
	result := collection.FindOneAndUpdate(context.TODO(), filter, update)
	return result.Err()
}

func CheckPermissionsForRole(accessRequestPath string, roleId int) error {
	role, err := getRoleById(roleId)
	if err != nil {
		return err
	}
	rolePaths := role.Paths
	for _, path := range rolePaths {
		if hasAccess(accessRequestPath, path.Path) {
			return nil
		}
	}
	return errors.New("You haven't permission to " + accessRequestPath + ".")
}

func getRoleById(roleId int) (*entity.UserRole, error) {
	collection := client.Database(SK_DB_NAME).Collection(ROLES_COLLECTION_NAME)
	filter := bson.M {
		"_id": roleId,
	}
	var role *entity.UserRole
	var err error
	err = collection.FindOne(context.TODO(), filter).Decode(&role)
	return role, err
}

func getRoleByName(roleName string) (*entity.UserRole, error) {
	collection := client.Database(SK_DB_NAME).Collection(ROLES_COLLECTION_NAME)
	filter := bson.M{
		"name": roleName,
	}
	var role *entity.UserRole
	var err error
	err = collection.FindOne(context.TODO(), filter).Decode(&role)
	return role, err
}

func checkRoleIsExist(roleId int) error {
	collection := client.Database(SK_DB_NAME).Collection(ROLES_COLLECTION_NAME)
	roleFilter := bson.D{
		{"_id", roleId},
	}
	singleResult := collection.FindOne(context.TODO(), roleFilter)
	return singleResult.Err()
}
