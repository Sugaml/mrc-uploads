package controller

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (ctl *Controller) handleFileupload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("image upload error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})

	}
	uniqueId := uuid.New()
	filename := strings.Replace(uniqueId.String(), "-", "", -1)
	fileExt := strings.Split(file.Filename, ".")[1]
	image := fmt.Sprintf("%s.%s", filename, fileExt)
	err = c.SaveFile(file, fmt.Sprintf("./uploads/%s", image))
	if err != nil {
		log.Println("image save error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}
	imageUrl := fmt.Sprintf(os.Getenv("BASE_URL")+"/uploads/%s", image)
	data := map[string]interface{}{
		"name":    image,
		"fileUrl": imageUrl,
		"header":  file.Header,
		"size":    file.Size,
	}
	return c.JSON(fiber.Map{"status": 201, "message": "Image uploaded successfully", "data": data})
}

func (ctl *Controller) handleDeleteImage(c *fiber.Ctx) error {
	// extract image name from params
	imageName := c.Params("imageName")

	// delete image from ./images
	err := os.Remove(fmt.Sprintf("./uploads/%s", imageName))
	if err != nil {
		log.Println(err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server Error", "data": nil})
	}

	return c.JSON(fiber.Map{"status": 201, "message": "Image deleted successfully", "data": nil})
}

func (ctl *Controller) handleMultipleFileupload(c *fiber.Ctx) error {

	// parse incomming image file
	file, err := c.FormFile("image")
	if err != nil {
		log.Println("image upload error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})

	}

	signature, err := c.FormFile("signature")
	if err != nil {
		log.Println("image upload error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})

	}

	// generate new uuid for image name
	uniqueId := uuid.New()

	// remove "- from imageName"

	filename := strings.Replace(uniqueId.String(), "-", "", -1)

	// extract image extension from original file filename

	fileExt := strings.Split(file.Filename, ".")[1]
	signatureExt := strings.Split(signature.Filename, ".")[1]

	// generate image from filename and extension
	image := fmt.Sprintf("%s.%s", filename, fileExt)
	sign := fmt.Sprintf("%s.%s", filename, signatureExt)
	// save image to ./images dir
	err = c.SaveFile(file, fmt.Sprintf("./uploads/%s", image))
	if err != nil {
		log.Println("image save error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}
	err = c.SaveFile(file, fmt.Sprintf("./uploads/%s", sign))
	if err != nil {
		log.Println("image save error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	// generate image url to serve to client using CDN
	signUrl := fmt.Sprintf(os.Getenv("BASE_URL")+"/uploads/%s", sign)
	imageUrl := fmt.Sprintf(os.Getenv("BASE_URL")+"/uploads/%s", image)

	// create meta data and send to client
	data := map[string]interface{}{

		"imageName": image,
		"imageUrl":  imageUrl,
		"signUrl":   signUrl,
		"header":    file.Header,
		"size":      file.Size,
	}
	return c.JSON(fiber.Map{"status": 201, "message": "Image uploaded successfully", "data": data})
}
