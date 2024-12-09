package handlers

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/maragudk/gomponents"
)

func HandleChatPost(c *fiber.Ctx) error {
	message := c.FormValue("message")
	if message == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Message is required")
	}

	// Add user message
	userMessage := gomponents.El("div",
		gomponents.Attr("class", "ai-message user"),
		gomponents.Text(message),
	)

	// TODO: Integrate with your preferred AI service
	// For now, just echo back a response
	aiResponse := "I received your message: " + message

	aiMessage := gomponents.El("div",
		gomponents.Attr("class", "ai-message"),
		gomponents.Text(aiResponse),
	)

	// Render both messages
	var buf bytes.Buffer
	if err := userMessage.Render(&buf); err != nil {
		return err
	}
	if err := aiMessage.Render(&buf); err != nil {
		return err
	}

	return c.SendString(buf.String())
}
