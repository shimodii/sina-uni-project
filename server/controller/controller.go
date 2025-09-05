package controller

import (
  "github.com/gofiber/fiber/v2"
  "fmt"
)

func Root (c *fiber.Ctx) error {
  return c.SendString("server is up...")
}

func Upload (c *fiber.Ctx) error {
	file, err := c.FormFile("file")
  fmt.Println(file)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "file is required",
		})
	}

	filePath := fmt.Sprintf("./uploads/%s", file.Filename)
	if err := c.SaveFile(file, filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to save file",
		})
	}

	return c.JSON(fiber.Map{
		"message": "file uploaded successfully",
		"path":    filePath,
	})
}
