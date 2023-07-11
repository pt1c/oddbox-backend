package main

import (
	"oddbox/src/database"
	"oddbox/src/models"
)

func main() {
	database.Connect()

	var itemTypes []models.ItemsType
	var itemType models.ItemsType

	itemType = models.ItemsType{
		Title: "Книги",
		Parameters: []*models.TypesParameter{
			{ Title: "ISBN", Description: "Номер книги" },
			{ Title: "Обложка", Description: "Твердая или мягкая" },
			{ Title: "Вес", Description: "Вес книги в граммах" },
		},
	}
	itemTypes = append(itemTypes, itemType)

	itemType = models.ItemsType{
		Title: "Лекарства",
		Parameters: []*models.TypesParameter{
			{ Title: "Срок годности", Description: "Срок годности" },
			{ Title: "Форма", Description: "Таблетки/сироп/свечи/ампулы/etc." },
		},
	}
	itemTypes = append(itemTypes, itemType)

	itemType = models.ItemsType{
		Title: "Резисторы",
		Parameters: []*models.TypesParameter{
			{ Title: "Корпус", Description: "DIP,etc" },
			{ Title: "Сопротивление", Description: "В ОМ-ах" },
			{ Title: "Допуск", Description: "в %" },
		},
	}
	itemTypes = append(itemTypes, itemType)

	database.DB.CreateInBatches(itemTypes, 3)

}