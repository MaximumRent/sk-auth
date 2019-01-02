package entity

import (
	"time"
)

// User's auth token, chain of this tokens can tell all user's login-logout history.
// LoginTime - time when user was logged in.
// LogoutTime - time when user log outed.
// AuthDevice - device which was used for auth. Can be 'Mobile' or 'Web' and unique identifier for device.
// JWTToken - auth token.
type AuthToken struct {
	LoginTime  time.Time   `bson:"loginTime"`
	LogoutTime time.Time   `bson:"logoutTime"`
	AuthDevice *AuthDevice `json:"authDevice" bson:"authDevice"`
	JWTToken   *jwt.Token  `json:"token" bson:"token"`
}
