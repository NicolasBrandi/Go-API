package main

import (
	"log"

	"github.com/NicolasBrandi/Go-API/database"
	"github.com/NicolasBrandi/Go-API/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error{
	return c.SendString("Welcome to the API")
}

func setUpRoutes(app *fiber.App){
	app.Get("/api",welcome)
	//Creating user endpoint
	app.Post("/api/users", routes.CreateUser)
}

func main(){
	//innit db
	database.ConnectDb()
	//innit fiber
	app := fiber.New()
	//Endpoints handler func
	setUpRoutes(app)


	log.Fatal(app.Listen(":3000"))
}

//'/Users/nicolasbrandi/go/bin/air' to call the module for autoreloading

