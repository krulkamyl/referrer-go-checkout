package auth

import (
	"referrer/app/database"
	"referrer/app/http/middlewares"
	"referrer/app/models"

	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Show(c *fiber.Ctx) error {

	id, _ := middlewares.GetUserId(c)

	var user models.User

	database.DB.Where("id = ?", id).First(&user)

	if strings.Contains(c.Path(), "/api/referrer") {
		referrer := models.Referrer(user)
		referrer.CalculateRevenue(database.DB)
	}

	return c.JSON(user)
}

func Update(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	id, _ := middlewares.GetUserId(c)

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}

	user.Id = id

	database.DB.Model(models.User{}).Updates(&user)

	return c.JSON(user)
}

func UpdatePassword(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Passwords do not match",
		})
	}

	id, _ := middlewares.GetUserId(c)

	user := models.User{}
	user.Id = id

	user.SetPassword(data["password"])

	database.DB.Model(models.User{}).Updates(&user)

	return c.JSON(fiber.Map{
		"message": "Passwords updated!",
	})
}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	isReferrer := strings.Contains(c.Path(), "/api/referrer")

	var scope string

	if isReferrer {
		scope = "referrer"
	} else {
		scope = "admin"
	}

	if !isReferrer && user.IsReferrer {
		c.Status(fiber.StatusUnauthorized)

		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	token, err := middlewares.GenerateJWT(user.Id, scope)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	now := time.Now()
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  now.Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Passwords do not match",
		})
	}

	user := models.User{
		FirstName:  data["first_name"],
		LastName:   data["last_name"],
		Email:      data["email"],
		IsReferrer: strings.Contains(c.Path(), "/api/refereer"),
	}
	user.SetPassword(data["password"])

	database.DB.Create(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {

	now := time.Now()

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  now.Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Logout success!",
	})
}
