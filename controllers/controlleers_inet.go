package controllers

import (
	"regexp"
	"strconv"
	"strings"

	"go-workshop/database"
	m "go-workshop/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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

func DogIDGreaterThan100(db *gorm.DB) *gorm.DB {
	return db.Where("dog_id > ?", 100)
}

// Read Dogs
func GetDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //delelete = null
	return c.Status(200).JSON(dogs)
}

// Read Dogs
func GetDog(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var dog []m.Dogs

	result := db.Find(&dog, "dog_id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&dog)
}

// Create Dogs
func AddDog(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn
	var dog m.Dogs

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&dog)
	return c.Status(201).JSON(dog)
}

// Update Dogs
func UpdateDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs
	id := c.Params("id")

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

// Delete Dog
func RemoveDog(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var dog m.Dogs

	result := db.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func GetDogsJson(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //10ตัว

	var dataResults []m.DogsRes
	var CountRed, CountGreen, CountPink, CountNo int
	for _, v := range dogs { //1 inet 112 //2 inet1 113
		typeStr := ""
		if v.DogID > 10 && v.DogID < 50 {
			typeStr = "red"
			CountRed++
		} else if v.DogID > 100 && v.DogID < 150 {
			typeStr = "green"
			CountGreen++
		} else if v.DogID > 200 && v.DogID < 250 {
			typeStr = "pink"
			CountPink++
		} else {
			typeStr = "no color"
			CountNo++
		}

		d := m.DogsRes{
			Name:  v.Name,  //inet1
			DogID: v.DogID, //113
			Type:  typeStr, //green
		}
		dataResults = append(dataResults, d)
		// sumAmount += v.Amount
	}

	r := m.ResultData{
		Count:      len(dogs), //หาผลรวม,
		Data:       dataResults,
		Name:       "golang-test",
		CountRed:   CountRed,
		CountGreen: CountGreen,
		CountPink:  CountPink,
		CountNo:    CountNo,
	}
	return c.Status(200).JSON(r)
}

func TestShowDelete(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs
	db.Unscoped().Where("deleted_at IS NOT NULL").Find(&dogs)
	return c.Status(200).JSON(dogs)
}

// Read Companies
func GetCompanies(c *fiber.Ctx) error {
	db := database.DBConn
	var companies []m.Company

	db.Find(&companies) //delelete = null
	return c.Status(200).JSON(companies)
}

// Read Company By Id With Create
func GetCompany(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var company []m.Company

	result := db.Find(&company, "company_id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&company)
}

// Create Company
func AddCompany(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn
	var company m.Company

	if err := c.BodyParser(&company); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&company)
	return c.Status(201).JSON(company)
}

// Update Company
func UpdateCompany(c *fiber.Ctx) error {
	db := database.DBConn
	var company m.Company
	id := c.Params("id")

	if err := c.BodyParser(&company); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&company)
	return c.Status(200).JSON(company)
}

// Delete Company
func RemoveCompany(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var company m.Company

	result := db.Delete(&company, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func GetDogsScope(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Scopes(GetScope).Find(&dogs) //delelete = null
	return c.Status(200).JSON(dogs)
}

func GetScope(db *gorm.DB) *gorm.DB {
	return db.Where("dog_id > ? && dog_id < ?", 50, 100)
}
