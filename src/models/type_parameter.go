package models

// Items_Parameter struct
type TypesParameter struct {
	Id       		uint   	`gorm:"primaryKey" json:"param_id"`
  Title    		string	`gorm:"not null" json:"param_title"`
	Description	string 	`json:"param_description"`

	Types []*ItemsType `gorm:"many2many:types_2_parameters;"`
}