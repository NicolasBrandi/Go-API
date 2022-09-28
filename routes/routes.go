package routes

import (
	"github.com/NicolasBrandi/Go-API/database"
	"github.com/NicolasBrandi/Go-API/models"
	"github.com/gofiber/fiber/v2"
)

//This is a serializer: struct that i am going to use to reconstruct the user with new data. It is != than models.User{}
type UserSer struct{
	ID 			uint 	`json:"id"`
	FirstName 	string	`json:"first_name"`
	LastName 	string	`json:"last_name"`
}

//Takes as argument an user from models and returns an User from routes 
func CreateResponseUser(userModel models.User) UserSer {
	return UserSer{ID: userModel.ID, FirstName: userModel.FirstName, LastName: userModel.LastName}
}

//ROUTES
//First endpoint
func CreateUser(c *fiber.Ctx) error{
	var user models.User

	if err := c.BodyParser(&user); err != nil{
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

//Second endpoint

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	//Take the slice users and .Find() looks for all the matching var into the struct. Gorm gives back all the matching users
	database.Database.Db.Find(&users)
	respUsers := []UserSer{}

	for _, user := range users{
		respUser := CreateResponseUser(user)
		respUsers = append(respUsers, respUser)
	}

	return c.Status(200).JSON(respUsers)
}