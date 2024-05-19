package controllers

import (
	"strconv"

	m "go-workshop/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func FiveDotOne(c *fiber.Ctx) error {
	stn := c.Params("number")
	num, err := strconv.Atoi(stn)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString("Error")
	}
	count := 1
	for i := num; i > 0; i-- {
		count *= i
	}
	result := c.Params("number") + "!=" + strconv.Itoa(count)
	return c.JSON(result)
}

func FiveDotTwo(c *fiber.Ctx) error {
	c.Query("tax_id") // "fenny"
	a := c.Query("tax_id")
	var Result string
	var text string
	for _, v := range a {
		Result = strconv.Itoa(int(v))
		text = text + " " + Result
	}
	return c.JSON(text)
}

func Six(c *fiber.Ctx) error {
	Data := new(m.Person)

	if err := c.BodyParser(Data); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("error")
	}

	validate := validator.New()
	if errors := validate.Struct(Data); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	}

	return c.JSON(Data)
}
