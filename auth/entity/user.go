package entity

import (
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/mgo.v2/bson"
	"sk-auth/auth/crypto"
	"sk-auth/validation"
	"time"
)

// System user which uses for authentication in all systems.
// Here, 'Mandatory' means this info is mandatory for user creating.
// Nickname - Mandatory. Unique. User's nickname for all systems.
// Email - Mandatory. Unique. User's email.
// Password - Mandatory. Hashed password.
// FirstName - Not mandatory. User's first name.
// SecondName - Not mandatory. User's second name.
// Gender - Not mandatory. User gender. Can be M or F. Not a "Sign up info".
// PhoneNumber - Not mandatory. Phone number should be configurated in settings and isn't a "Sign up info".
// CreatedTime - Not mandatory. Timestamp when user was created.
// AuthTokens - Not mandatory. "History". All tokens that was used by this user. Don't used for json serialization.
// Roles - Mandatory. We don't need to put this field to json, so we have only bson mapping.
type User struct {
	Id             bson.ObjectId `json:"id" bson:"_id"`
	Nickname       string        `json:"nickname" bson:"nickname" validate:"required"`
	Email          string        `json:"email" bson:"email" validate:"required,email"`
	Password       string        `json:"password" bson:"password" validate:"required"`
	FirstName      string        `json:"firstName" bson:"firstName"`
	LastName       string        `json:"lastName" bson:"lastName"`
	Gender         string        `json:"gender" bson:"gender"`
	PhoneNumber    string        `json:"phoneNumber" bson:"phoneNumber"`
	CreatedTime    time.Time     `json:"createdTime" bson:"createdTime"`
	AuthTokens     []*AuthToken  `bson:"tokens"`
	Roles          []*UserRole   `bson:"roles"`
	selfValidation validation.SelfValidatable
}

func (self *User) SelfValidate() error {
	err := validator.New().Struct(self)
	if err == nil {
		self.Password, err = crypto.EncryptPassword(self.Password)
	}
	return err
}

// Factory function for User entity
func CreateUser() *User {
	user := new(User)
	user.Id = bson.NewObjectId()
	//user.Nickname = nickname
	//user.Email = email
	user.CreatedTime = time.Now()
	//encryptedPassword, err := crypto.EncryptPassword(password)
	// If we has some problems with encryption
	// set undefined value for password
	//if err == nil {
	//	user.Password = encryptedPassword
	//} else {
	//	user.Password = _UNDEFINED_PASSWORD
	//}
	//user.Roles = append(user.Roles, USER_ROLE)
	return user
}
