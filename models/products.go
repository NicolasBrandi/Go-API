package models

type Product struct{
	ID 				uint 	`json:"id" gorm:"primaryKey"`
	Name 			string	`json:"name"`
	SerialNumber	int		`json:"serial_number"`
}