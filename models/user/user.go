package user

import "time"

type User struct {
	Id        int64
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "user.users"
}
