package models

// Item struct
type Item struct {
	Id       	uint   	`gorm:"primaryKey" json:"item_id"`
  Title    	string	`gorm:"not null" json:"item_title"`
	UserId		uint		`json:"user_id"`
	TypeId		uint		`json:"item_type"`
	Qty			 	uint		`gorm:"not null" json:"qty"`

	Type			ItemsType `gorm:"foreignKey:TypeId"`
}