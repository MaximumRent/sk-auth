package mongo

import (
	"context"
	"errors"
	"gopkg.in/mgo.v2/bson"
	"log"
	"sk-auth/auth/crypto"
	"sk-auth/auth/entity"
	"time"
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
	filter := bson.M{
		"$or": []bson.M{
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
	filter := bson.M{
		"$or": []bson.M{
			{"email": user.Email},
			{"nickname": user.Nickname},
		},
	}
	// Here we don't update email, because it non-updatable user-info, email must be static
	// Password also don't updated from here
	update := bson.M{
		"$set": bson.M{
			"firstName":   user.FirstName,
			"lastName":    user.LastName,
			"nickname":    user.Nickname,
			"gender":      user.Gender,
			"phoneNumber": user.PhoneNumber,
		},
	}
	result := collection.FindOneAndUpdate(context.TODO(), filter, update)
	return result.Err()
}

func ValidateAuthToken(email, nickname, token string) error {
	collection := client.Database(SK_DB_NAME).Collection(USER_COLLECTION_NAME)
	filter := bson.M{
		"$or": []bson.M{
			{"email": email},
			{"nickname": nickname},
		},
	}
	var user = &entity.User{}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return err
	}

	hasValidToken := false
	for _, authToken := range user.AuthTokens {
		if authToken.Token == token && !authToken.LogoutTime.After(time.Date(2, 2, 2, 0, 0, 0, 0, time.UTC)) {
			hasValidToken = true
		}
	}
	if hasValidToken {
		return nil
	} else {
		return errors.New("You haven't valid auth token")
	}
}

func Logout(email, nickname, authDevice, token string) error {
	collection := client.Database(SK_DB_NAME).Collection(USER_COLLECTION_NAME)
	filter := bson.M{
		"$or": []bson.M{
			{"email": email},
			{"nickname": nickname},
		},
		"tokens.logoutTime": bson.M{
			"$lt": time.Date(2018, 2, 2, 0, 0, 0, 0, time.UTC),
		},
		"tokens.authDevice": bson.M{
			"$eq": authDevice,
		},
	}
	user := new(entity.User)
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return err
	}

	for _, authToken := range user.AuthTokens {
		if authToken.Token == token && authToken.AuthDevice == authDevice && !authToken.LogoutTime.After(time.Date(2, 2, 2, 0, 0, 0, 0, time.UTC)) {
			authToken.LogoutTime = time.Now()
		}
	}
	update := bson.M{
		"$set": bson.M{
			"tokens": user.AuthTokens,
		},
	}
	result := collection.FindOneAndUpdate(context.TODO(), filter, update)
	//update := bson.M{
	//	"$set": bson.M{
	//		"logoutTime": time.Now(),
	//	},
	//}
	////opt := &options.FindOneAndUpdateOptions{ArrayFilters: &options.ArrayFilters{Filters: []interface{}()},}
	//result := collection.FindOneAndUpdate(context.TODO(), filter, update)
	return result.Err()
}

func AddAuthToken(email string, token entity.AuthToken) error {
	collection := client.Database(SK_DB_NAME).Collection(USER_COLLECTION_NAME)
	filter := bson.M{
		"email": email,
	}
	update := bson.M{
		"$push": bson.M{
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
