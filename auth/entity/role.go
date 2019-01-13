package entity

type UserRole struct {
	Name        string `json:"name" bson:"name"`
	Code        int64  `json:"code" bson:"code"`
	IsDeletable bool   `json:"isDeletable" bson:"isDeletable"`
}

// DEFINED USER ROLES
var (
	ADMIN_ROLE = &UserRole{
		Name:        "Admin",
		Code:        0,
		IsDeletable: false,
	}
	USER_ROLE = &UserRole{
		Name:        "User",
		Code:        1,
		IsDeletable: false,
	}
	COMPANY_OWNER_ROLE = &UserRole{
		Name:        "Company Owner",
		Code:        2,
		IsDeletable: false,
	}
	COMPANY_MANAGER_ROLE = &UserRole{
		Name:        "Company Manager",
		Code:        3,
		IsDeletable: false,
	}
)
