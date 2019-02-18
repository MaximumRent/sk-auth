package mongo

import (
	"context"
	"gopkg.in/mgo.v2/bson"
	"sk-auth/auth/entity"
)

func AddPathToRole(roleName string, path string) error {
	collection := client.Database(SK_DB_NAME).Collection(ROLES_COLLECTION_NAME)
	filter := bson.M{
		"name": roleName,
	}
	update := bson.M{
		"$push": bson.M{
			"path": path,
		},
	}
	result := collection.FindOneAndUpdate(context.TODO(), filter, update)
	return result.Err()
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
