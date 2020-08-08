package model

type UsersUser struct {
	*Model
	Username string `json:"username"`
}

func (u UsersUser) TableName() string {
	return "users_user"
}
