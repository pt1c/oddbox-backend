package models

// Item_Type struct
type Item_Type struct {
	ID       	uint   	`gorm:"primaryKey" json:"type_id"`
  Title    	string	`gorm:"not null" json:"type_title"`
}