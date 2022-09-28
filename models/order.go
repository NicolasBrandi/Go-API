package models

type Order struct{
	ID 			 uint		`json:"id" gorm:"primaryKey"`
	//foraign key from users
	UserRefer int		`json:"user_id"`
	User 	 User	`gorm:"foreignKey: UserRefer"`
	//foreign key from products
	ProductRefer int		`json:"product_id"`
	Product 	 Product	`gorm:"foreignKey: ProductRefer"`
}