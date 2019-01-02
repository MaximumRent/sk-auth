package entity

import (
	"time"
)

// System user which uses for authentication in all systems.
// Here, 'Mandatory' means this info is mandatory for user creating.
// Nickname - Mandatory. Unique. User's nickname for all systems.
// Email - Mandatory. Unique. User's email.
// Password - Mandatory. Hashed password.
// FirstName - Mandatory. User's first name.
// SecondName - Mandatory. User's second name.
// Gender - Not mandatory. User gender. Can be M or F. Not a "Sign up info".
// PhoneNumber - Not mandatory. Phone number should be configurated in settings and isn't a "Sign up info".
// CreatedTime - Not mandatory. Timestamp when user was created.
// AuthTokens - Not mandatory. "History". All tokens that was used by this user. Don't used for json serialization.
// Role - Mandatory. We don't need to put this field to json, so we have only bson mapping.
type User struct {
	Id          int64        `json:"id" bson:"_id"`
	Nickname    string       `json:"nickname" bson:"nickname"`
	Email       string       `json:"email" bson:"email"`
	Password    string       `json:"password" bson:"password"`
	FirstName   string       `json:"firstName" bson:"firstName"`
	SecondName  string       `json:"secondName" bson:"secondName"`
	Gender      string       `json:"gender" bson:"gender"`
	PhoneNumber string       `json:"phoneNumber" bson:"phoneNumber"`
	CreatedTime time.Time    `json:"createdTime" bson:"createdTime"`
	AuthTokens  []*AuthToken `bson:"authTokens"`
	Role        *UserRole    `bson:"role"`
}

// Factory function for User entity
func CreateUser(nickname, email, password, firstName, secondName string) {
	user := new(User)
	user.Nickname = nickname
	user.Email = email
	// Use hashing
	//user.password =
	user.FirstName = firstName
	user.SecondName = secondName
	user.CreatedTime = time.Now()
	// Create default user role
	// user.Role =
}
