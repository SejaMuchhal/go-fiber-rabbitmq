package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/streadway/amqp"
)

func main() {
	// Define RabbitMQ server URL.
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")

	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}

	defer connectRabbitMQ.Close()

	// Opening a channel to our RabbitMQ instance over the connection we have already established.
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	// Declare Queues that we can publish and subscribe to

	_, err = channelRabbitMQ.QueueDeclare(
		"QueueService1",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	// Create a new Fiber instance
	app := fiber.New()

	// Add middleware
	app.Use(
		logger.New(), // add simple logger
	)

	// Add route
	app.Get("/send", func(c *fiber.Ctx) error {
		// Create a message to publish
		message := amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(c.Query("msg")),
		}

		// Attempt to publish a message to the queue
		if err := channelRabbitMQ.Publish(
			"",
			"QueueService1",
			false,
			false,
			message,
		); err != nil {
			return err
		}

		return nil
	})

	// Start Fiber API server
	log.Fatal(app.Listen(":3000"))

}