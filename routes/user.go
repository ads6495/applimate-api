package routes

import (
	"github.com/ads6495/applimate-api/database"
	"github.com/ads6495/applimate-api/models"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	// not a 1 : 1 to a user model, this is a serializer for incoming requests
	ID         uint   `json:"id"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
}

// returns a User Struct
func CreateResponseUser(userModel models.User) User {
	return User{ID: userModel.ID, First_Name: *userModel.FirstName, Last_Name: *userModel.LastName}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}
