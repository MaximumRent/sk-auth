package mongo

import (
	"context"
	"log"
	"sk-auth/auth/crypto"
	"sk-auth/auth/entity"
)

func CreateUser(email, password string) error {
	user := entity.CreateUser()
	user.Password, _ = crypto.EncryptPassword(password)
	user.Email = email
	collection := client.Database(SK_DB_NAME).Collection(USER_COLLECTION_NAME)
	result, err := collection.InsertOne(context.TODO(), user)
	log.Printf("Was created new user with id: %s", result.InsertedID)
	return err
}

func AddRightsToRole(roleName string, rights []string) error {
	collection := client.Database(SK_DB_NAME).Collection(ROLES_COLLECTION_NAME)

}
