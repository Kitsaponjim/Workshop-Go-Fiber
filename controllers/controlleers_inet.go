package controllers

import (
	"regexp"
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
	result := 1
	for i := num; i > 0; i-- {
		result *= i
	}
	Ans := c.Params("number") + "!=" + strconv.Itoa(result)
	return c.JSON(Ans)
}

func FiveDotTwo(c *fiber.Ctx) error {
	c.Query("tax_id") // "fenny"
	t := c.Query("tax_id")
	var Result string
	var text string
	for _, v := range t {
		Result = strconv.Itoa(int(v))
		text += " " + Result
	}
	return c.JSON(text)
}

func TestParams(c *fiber.Ctx) error {
	return c.SendString("TestParams")
}

func Six(c *fiber.Ctx) error {
	Data := new(m.Person)

	if err := c.BodyParser(Data); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("error")
	}

	validate := validator.New()
	CheckUsername, _ := regexp.MatchString(`^[a-zA-Z0-9_-]+$`, Data.Username)
	if !CheckUsername {
		return c.Status(fiber.StatusBadRequest).SendString("ใช้อักษรภาษาอังกฤษ (a-z), (A-Z), ตัวเลข(0-9) และเครื่องหมาย (_), (-)เท่านั้น")
	}

	CheckWebsite, _ := regexp.MatchString(`^[a-z0-9-]+$`, Data.NameWebsite)
	if !CheckWebsite {
		return c.Status(fiber.StatusBadRequest).SendString("2 - 30 ตัวอักษรต้องเป็นตัวอักษรภาษาอังกฤษตัวเล็ก(a-z) ตัวเลข(0-9) ห้ามใช้เครื่องหมายอักขระพิเศษยกเว้นเครื่อหมายขีด (-) ห้ามเว้นวรรคและห้ามใช้ภาษาไทย")
	}

	if errors := validate.Struct(Data); errors != nil {
		return c.Status(fiber.StatusBadRequest).SendString("คุณใส่ข้อมูลไม่ตรงตามเงื่อนไข")
	}

	return c.JSON(Data)
}
