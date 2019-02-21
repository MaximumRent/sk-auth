package mongo

import (
	"context"
	"errors"
	"gopkg.in/mgo.v2/bson"
	"log"
	"sk-auth/auth/crypto"
	"sk-auth/auth/entity"
)

func CreateUser(user entity.User) error {
	if user.Nickname == "" {
		return errors.New("User nickname can't be empty!")
	}
	collection := client.Database(SK_DB_NAME).Collection(USER_COLLECTION_NAME)
	userRole, err := getRoleByName(entity.USER_ROLE_NAME)
	if err != nil {
		return err
	}
	shortUserRole := &entity.ShortUserRole{Id: userRole.Id}
	user.Roles = append(user.Roles, shortUserRole)
	user.Gender = entity.UNDEFINED_GENDER
	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	log.Printf("Was created new user with id: %s\n", result.InsertedID)
	return err
}

func LoginUser(loginUserInfo entity.LoginUserInfo) (error, *entity.AuthToken, *entity.User) {
	collection := client.Database(SK_DB_NAME).Collection(USER_COLLECTION_NAME)
	filter := bson.M {
		"$or": []bson.M {
			{"email": loginUserInfo.Login},
			{"nickname": loginUserInfo.Login},
		},
	}
	user := new(entity.User)
	err := collection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		return err, nil, nil
	}
	err = crypto.ComparePasswords(loginUserInfo.Password, user.Password)

	if err != nil {
		return err, nil, nil
	}
	token := entity.GenerateAuthToken(loginUserInfo.AuthDevice)
	err = AddAuthToken(user.Email, *token)
	return err, token, user
}

func UpdateUserInfo(user entity.User) error {
	collection := client.Database(SK_DB_NAME).Collection(USER_COLLECTION_NAME)
	filter := bson.D{
		{"email", user.Email},
	}
	update := bson.D{
		{"email", user.Email},
		{"password", user.Password},
		{"firstName", user.FirstName},
		{"lastName", user.LastName},
		{"nickname", user.Nickname},
		{"gender", user.Gender},
		{"phoneNumber", user.PhoneNumber},
	}
	result := collection.FindOneAndUpdate(context.TODO(), filter, update)
	return result.Err()
}

func ValidateAuthToken(email, token string) error {
	collection := client.Database(SK_DB_NAME).Collection(USER_COLLECTION_NAME)
	filter := bson.D{
		{"email", email},
		{"tokens", bson.D{
			{"token", token},
		}},
	}
	result := collection.FindOne(context.TODO(), filter)
	return result.Err()
}

func AddAuthToken(email string, token entity.AuthToken) error {
	collection := client.Database(SK_DB_NAME).Collection(USER_COLLECTION_NAME)
	filter := bson.M {
		"email": email,
	}
	update := bson.M {
		"$push": bson.M {
			"tokens": token,
		},
	}
	result := collection.FindOneAndUpdate(context.TODO(), filter, update)
	return result.Err()
}

func UpdateAuthToken(email string, token entity.AuthToken) error {
	collection := client.Database(SK_DB_NAME).Collection(USER_COLLECTION_NAME)
	filter := bson.D{
		{"email", email},
		{"tokens", bson.D{
			{"token", token.Token},
		}},
	}
	update := bson.D{
		{"tokens", bson.D{
			{"loginTime", token.LoginTime},
			{"logoutTime", token.LogoutTime},
			{"authDevice", token.AuthDevice},
		}},
	}
	result := collection.FindOneAndUpdate(context.TODO(), filter, update)
	return result.Err()
}

func AddRoleToUser(email string, roleId int) error {
	collection := client.Database(SK_DB_NAME).Collection(USER_COLLECTION_NAME)
	userFilter := bson.D{
		{"email", email},
	}
	err := checkRoleIsExist(roleId)
	if err != nil {
		return err
	}
	update := bson.D{
		{"$push", bson.D{
			{"roles", roleId},
		}},
	}
	result := collection.FindOneAndUpdate(context.TODO(), userFilter, update)
	return result.Err()
}
