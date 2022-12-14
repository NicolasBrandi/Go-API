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
	//Creating user endpoints for user
	app.Get("/api",welcome)
	app.Get("/api/users", routes.GetUsers)
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users/:id", routes.GetUser)
	app.Delete("/api/users/:id", routes.DeleteUser)
	//Creating user endpoints for products
	app.Post("/api/products", routes.CreateProduct)
	app.Get("/api/products", routes.GetProducts)
	app.Get("/api/products:id", routes.GetProduct)
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

