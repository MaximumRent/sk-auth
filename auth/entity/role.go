package entity

const (
	ADMIN_ROLE_ID      = 1
	USER_ROLE_ID       = 2
	COMPANY_OWNER_ID   = 3
	COMPANY_MANAGER_ID = 4
)

const (
	ADMIN_ROLE_NAME      = "admin"
	USER_ROLE_NAME       = "user"
	COMPANY_OWNER_NAME   = "companyOwner"
	COMPANY_MANAGER_NAME = "companyManager"
)

type ShortUserRole struct {
	Id int `bson:"role_id"`
}

type UserRole struct {
	Id          int         `bson:"_id"`
	Name        string      `json:"name" bson:"name"`
	Code        int64       `json:"code" bson:"code"`
	IsDeletable bool        `json:"isDeletable" bson:"isDeletable"`
	Paths       []*RolePath `bson:"paths"`
}

type RolePath struct {
	Path string `bson:"path"`
}
