package controller

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/wayanputri/blog/database"
	"github.com/wayanputri/blog/model"
)

// Blog List
func BlogList(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Blog list",
	}

	time.Sleep(time.Millisecond * 1500)
	db := database.DBConn

	var records []model.Blog
	if result := db.Find(&records); result.Error != nil {
		context["statusText"] = ""
		context["msg"] = "failed read data in database"
		c.Status(500)
		return c.JSON(context)
	}
	context["blog_records"] = records
	c.Status(200)
	return c.JSON(context)

}

func BlogDetail(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "",
		"msg":        "",
	}

	id := c.Params("id")
	var record model.Blog

	database.DBConn.First(&record, id)
	if record.ID == 0 {
		log.Println("Record not found.")
		context["msg"] = "Record not found."
		c.Status(404)
		return c.JSON(context)
	}

	context["statusText"] = "ok"
	context["msg"] = "Read Detail"
	context["record"] = record
	c.Status(200)
	return c.JSON(context)
}

// Add a blog into database
func BlogCreate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Add aBlog",
	}

	record := new(model.Blog)
	if err := c.BodyParser(record); err != nil {
		log.Println("Error in parsing request.")
		context["statusText"] = ""
		context["msg"] = "Failed body Parse"
		c.Status(400)
		return c.JSON(context)
	}

	file, err := c.FormFile("file")
	if err != nil {
		log.Println("Error in file upload.", err)
	}
	if file.Size > 0 {
		filename := "./static/uploads/" + file.Filename
		if err := c.SaveFile(file, filename); err != nil {
			log.Println("Error in file uploading...", err)
		}
		record.Image = filename
	}
	result := database.DBConn.Create(&record)
	if result.Error != nil {
		log.Fatal("Error in saving data.")
		context["statusText"] = ""
		context["msg"] = "Failed create data in database"
		c.Status(500)
		return c.JSON(context)
	}
	context["msg"] = "record is saved successfully"
	context["data"] = record

	c.Status(201)
	return c.JSON(context)
}

// Update a Blog
func BlogUpdate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Update a Blog",
	}
	id := c.Params("id")
	var record model.Blog

	database.DBConn.First(&record, id)
	if record.ID == 0 {

		context["statusText"] = ""
		context["msg"] = "Record not Found."
		c.Status(400)
		log.Println("Record not Found.")

		return c.JSON(context)
	}
	if err := c.BodyParser(&record); err != nil {
		context["statusText"] = ""
		context["msg"] = "Error in parsing request."
		c.Status(400)
		log.Println("Error in parsing request.")
		return c.JSON(context)
	}
	result := database.DBConn.Save(record)
	if result.Error != nil {
		context["statusText"] = ""
		context["msg"] = "Error in saving data."
		c.Status(500)
		log.Println("Error in saving data.")
		return c.JSON(context)
	}
	context["msg"] = "Record update successfully"
	context["data"] = record
	c.Status(200)
	return c.JSON(context)
}

// delete a Blog
func BlogDelete(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "",
		"msg":        "",
	}

	id := c.Params("id")
	var record model.Blog

	database.DBConn.First(&record, id)
	if record.ID == 0 {
		log.Println("Record not found.")
		context["msg"] = "Record not found."
		c.Status(400)
		return c.JSON(context)
	}
	result := database.DBConn.Delete(record)
	if result.Error != nil {
		log.Println("Failed delete data")
		context["msg"] = "Record not found."
		c.Status(500)
		return c.JSON(context)
	}
	context["statusText"] = "ok"
	context["msg"] = "Record delete successfully"
	c.Status(200)
	return c.JSON(context)
}
