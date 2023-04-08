package routes

import (
	"go-crud-fiber/intializers"
	"go-crud-fiber/models"

	_ "github.com/arsmn/fiber-swagger/v2/example/docs"
	"github.com/gofiber/fiber/v2"
)

// AddPost godoc
// @Summary Add a new post
// @Description Add a new post to the database

// @Accept  json
// @Produce  json
// @Tags models.PostModel
// @Param models.PostModel
// @Success 201 models.PostModel
// @Success 400 {object}

// @Router /api/addpost/ [post]
func AddPost(c *fiber.Ctx) error {
	post := new(models.PostModel)
	if err := c.BodyParser(post); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "title or body is empty", "data": nil})
	}
	if post.Title == "" || post.Body == "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "title or body is empty", "data": nil})
	}

	intializers.DB.Create(&post)

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Post successfully created", "data": post})

}

func AllPosts(c *fiber.Ctx) error {
	var posts []models.PostModel
	intializers.DB.Find(&posts)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Post fetched successfully", "data": posts})
}

// Post represents a post in the database
func Post(c *fiber.Ctx) error {
	// Parse the request body into a struct that represents the request
	var req struct {
		ID uint `json:"id"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No id found", "data": nil})
	}
	id := req.ID
	db := intializers.DB
	var post models.PostModel
	db.Find(&post, id)
	if post.Title == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No post found with ID", "data": nil})

	}
	return c.JSON(fiber.Map{"status": "success", "message": "post found", "data": post})

}

// DeletePost is handler/controller which deletes Post in the Database

func DeletePost(c *fiber.Ctx) error {
	// Parse the request body into a struct that represents the request
	var req struct {
		ID uint `json:"id"`
	}
	var post models.PostModel
	if err := c.BodyParser(&req); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No post found with ID", "data": nil})
	}
	db := intializers.DB

	db.Find(&post, req.ID)
	if post.Title == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No post found with ID", "data": nil})

	}
	intializers.DB.Delete(&post, req.ID)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "post deleted", "data": nil})
}

// UpdatePost is handler/controller which updates Post in the Database

func UpdatePost(c *fiber.Ctx) error {
	var post models.PostModel
	if err := c.BodyParser(&post); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	intializers.DB.First(&post, post.ID)

	intializers.DB.Save(&post)

	return c.Status(200).JSON(post)
}
