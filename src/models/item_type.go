package models

// ItemsType struct
type ItemsType struct {
	Id       	uint   	`gorm:"primaryKey" json:"type_id"`
  Title    	string	`gorm:"not null" json:"type_title"`

	Parameters []*TypesParameter `gorm:"many2many:types_2_parameters;"`
}