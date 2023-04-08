package main

import (
	"go-crud-fiber/intializers"
	"go-crud-fiber/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	intializers.LoadEnvironmentVariables()
	intializers.ConnectDB()
}

func main() {

	app := fiber.New()

	setUpRoutes(app)

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
	err := app.Listen(":3000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
func setUpRoutes(app *fiber.App) {

	app.Get("/api/allposts", routes.AllPosts)
	app.Post("/api/addpost", routes.AddPost)
	app.Post("/api/post/", routes.Post)
	app.Put("/api/updatepost", routes.UpdatePost)
	app.Post("/api/deletepost", routes.DeletePost)
}
