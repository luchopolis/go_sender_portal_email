package main

import (
	"fmt"
	"log"
	"os"

	"net/smtp"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type MessageInfo struct {
	To      string `json:"to" xml:"to" form:"to"`
	Message string `json:"message" xml:"message" form:"message"`
}

func mailSender(to string, content string) {

	emailKey := os.Getenv("EMAIL_KEY")
	emailSender := os.Getenv("EMAIL_SENDER")

	auth := smtp.PlainAuth("", emailSender, emailKey, "smtp.gmail.com")

	toMsg := []string{to}

	msg := []byte(content)

	errMail := smtp.SendMail("smtp.gmail.com:587", auth, emailSender, toMsg, msg)

	if errMail != nil {

		log.Fatal(errMail)

	}
}

func main() {
	fmt.Println("Running")

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	app := fiber.New()

	app.Post("/api/send-mail", func(c *fiber.Ctx) error {

		message := new(MessageInfo)

		if err := c.BodyParser(message); err != nil {
			return err
		}

		fmt.Println(message.Message)
		fmt.Println(message.To)

		mailSender(message.To, message.Message)

		return c.SendString("Excelente")
	})

	// port, errorParsing := strconv.Atoi(os.Getenv("PORT"))

	// if errorParsing != nil {
	// 	log.Fatal("Error Converting")
	// }

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))

}
