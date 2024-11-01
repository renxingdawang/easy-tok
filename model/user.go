package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"index:idx_username,unique;type:varchar(40);not null" json:"username"`
	Password string `gorm:"type:varchar(256);not null" json:"password"`
	//FavoriteVideos []Video `gorm:"many2many:user_favorite_videos" json:"favorite_videos"`
	FollowingCount int `gorm:"default:0" json:"following_count"`
	FollowerCount  int `gorm:"default:0" json:"follower_count"`
}

func (u *User) TableName() string {
	return "users"
}
