package entity

import (
	"github.com/google/uuid"
	"time"
)

// User's auth token, chain of this tokens can tell all user's login-logout history.
// LoginTime - time when user was logged in.
// LogoutTime - time when user log outed.
// AuthDevice - device which was used for auth. Can be 'Mobile' or 'Web' and unique identifier for device.
// Token - auth token.
type AuthToken struct {
	LoginTime  time.Time   	`bson:"loginTime"`
	LogoutTime time.Time   	`bson:"logoutTime"`
	AuthDevice string 		`json:"authDevice" bson:"authDevice"`
	Token      string      	`json:"token" bson:"token"`
}

func GenerateAuthToken(authDevice string) *AuthToken {
	token := new(AuthToken)
	token.AuthDevice = authDevice
	token.LoginTime = time.Now()
	token.Token = uuid.New().String()
	return token
}
