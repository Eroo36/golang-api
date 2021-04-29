package models

import (
	"time"

	"github.com/goonode/mogo"
)

//User struct is to handle user data
type User struct {
	mogo.DocumentModel `bson:",inline" coll:"users"`
	Email              string `idx:"{email},unique" json:"email" binding:"required"`
	Password           string `json:"password" binding:"required"`
	Name               string `json:"name" binding:"required"`
	CreatedAt          *time.Time
	UpdatedAt          *time.Time
	VerifiedAt         *time.Time
}

func UserInit() {
	mogo.ModelRegistry.Register(User{})
}
