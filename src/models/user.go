package models

type User struct {
	Id       uint   `json:"id"`
	Username string `gorm:"uniqueIndex;not null;type:varchar(64)" json:"username" validate:"gte=4"`
	Email    string `gorm:"uniqueIndex;not null;type:varchar(64)" json:"email" validate:"email"`
	Password string `gorm:"not null" json:"-" validate:"gte=8"`
	Items    []Item `gorm:"foreignKey:UserId" json:"items"`
}