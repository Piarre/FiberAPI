package routes

import (
	"SQLFiberApi/database"
	"SQLFiberApi/models"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID 		 uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func FindUserById(id int, user *models.User) error {
	database.Database.DB.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}

func CreateResponseUser(userModel models.User) User {
	return User{ID: userModel.ID, FirstName: userModel.FirstName, LastName: userModel.LastName,}
}

func CreateUser(ctx *fiber.Ctx) error {
	var user models.User
	
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	database.Database.DB.Create(&user)
	responseUser := CreateResponseUser(user)

	return ctx.Status(200).JSON(responseUser)
}

func GetUsers(ctx *fiber.Ctx) error {
	users := []models.User{}

	database.Database.DB.Find(&users)
	responseUsers := []User{}
	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return ctx.Status(200).JSON(responseUsers)
}

func GetUserById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	var user models.User

	if err != nil {
		return ctx.Status(400).JSON("Please, ensure that :id is an integer!")
	}

	if err := FindUserById(id, &user); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	responseUser := CreateResponseUser(user)
	return ctx.Status(200).JSON(responseUser)
}

func UpdateUserById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	var user models.User

	if err != nil {
		return ctx.Status(400).JSON("Please, ensure that :id is an integer!")
	}

	if err := FindUserById(id, &user); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	type UserToUpdate struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var dataToUpdate UserToUpdate

	if err := ctx.BodyParser(&dataToUpdate); err != nil {
		return ctx.Status(500).JSON(err.Error())
	}

	user.FirstName = dataToUpdate.FirstName
	user.LastName  = dataToUpdate.LastName

	database.Database.DB.Save(&user)

	responseUser := CreateResponseUser(user)
	return ctx.Status(200).JSON(responseUser)
}

func DeleteUserById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	var user models.User

	if err != nil {
		return ctx.Status(400).JSON("Please, ensure that :id is an integer!")
	}

	if err := FindUserById(id, &user); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	if err := database.Database.DB.Delete(&user).Error; err != nil {
		return ctx.Status(404).JSON(err.Error())
	}

	return ctx.Status(200).SendString("User successfully deleted")
}