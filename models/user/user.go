package user

import "time"

type User struct {
	Id             int64      `gorm:"column:id;primaryKey;autoIncrement"`
	Email          string     `gorm:"column:email;unique;not null"`
	Password       string     `gorm:"column:password;not null"`
	Username       string     `gorm:"column:username;unique;not null"`
	Name           *string    `gorm:"column:name"`
	Phone          *string    `gorm:"column:phone"`
	Birthdate      *time.Time `gorm:"column:birthdate"`
	Avatar         *string    `gorm:"column:avatar"`
	Biodata        *string    `gorm:"column:biodata"`
	City           *string    `gorm:"column:city"`
	Country        *string    `gorm:"column:country"`
	FollowersTotal int        `gorm:"column:followers_total;default:0"`
	FollowingTotal int        `gorm:"column:following_total;default:0"`
	PrivateAccount bool       `gorm:"column:private_account;default:false"`
	CreatedAt      time.Time  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      time.Time  `gorm:"column:updated_at;autoUpdateTime"`
}

func (User) TableName() string {
	return "user.users"
}
