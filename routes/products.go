package routes

import (
	"errors"

	"github.com/NicolasBrandi/Go-API/database"
	"github.com/NicolasBrandi/Go-API/models"
	"github.com/gofiber/fiber/v2"
)

type ProductSer struct{
	Name 			string	`json:"name"`
	SerialNumber	int		`json:"serial_number"`
}

func CreateResponseProduct(productModel models.Product) ProductSer{
	return ProductSer{Name: productModel.Name, SerialNumber: productModel.SerialNumber}
}

func CreateProduct(c *fiber.Ctx) error{
	var product models.Product

	if err := c.BodyParser(&product); err != nil{
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&product)
	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}

func GetProducts(c *fiber.Ctx) error{
	products := []models.Product{}
	//Finds all the products that exists and match with the structure of []models.product{}. So now products is an slice pointing to all the products that match
	database.Database.Db.Find(&products)
	//Var that i am going to return
	productsResponse := []ProductSer{}

	for _,prod := range products{
		prodUnitResponse := CreateResponseProduct(prod)
		productsResponse = append(productsResponse, prodUnitResponse)
	}

	return c.Status(200).JSON(productsResponse)

}

func findProduct(id int, prod *models.Product) error{
	database.Database.Db.Find(&prod, "id = ?", id)
	//TODO: deal with err
	if prod.ID == 0 {
		return errors.New("ID not found")
	}
	return nil
}

func GetProduct(c *fiber.Ctx) error{
	id, err := c.ParamsInt("id")

	product := models.Product{}

	//Handling the err comming from .ParamsInt
	if err != nil {
		return c.Status(400).JSON("Enter an int")
	}

	//calling findProduct() here parse the user i am looking for into product
	if err := findProduct(id, &product); err != nil{
		return c.Status(400).JSON(err.Error())
	}

	//now that i have what i am looking for into product, i just have to create it with the serializer

	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)
}

func DeleteProduct(c *fiber.Ctx) error{
	id, err := c.ParamsInt("id")

	product := models.Product{}

	//Handling the err comming from .ParamsInt
	if err != nil {
		return c.Status(400).JSON("Enter an int")
	}

	//calling findProduct() here parse the user i am looking for into product
	if err := findProduct(id, &product); err != nil{
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&product).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.SendString("Product deleted")
}