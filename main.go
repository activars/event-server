package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/net/context"
	"log"
	"cloud.google.com/go/pubsub"
	"io/ioutil"
	"net/http"
	"github.com/labstack/echo/engine/standard"
	"fmt"
)

const(
	topicName = "events"
	projectId = "development-50atoms"
)

func main() {
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("creating pubsub client: %v", err)
	}
	topic := client.Topic(topicName)

	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Post("/event", func(c echo.Context) error {

		msg, err := ioutil.ReadAll(c.Request().Body())
		if err != nil {
			return 	c.JSON(http.StatusAccepted, map[string]string {
				"error": "invalid message",
			})
		} else {
			msgID, err := topic.Publish(ctx, &pubsub.Message{
				Data: msg,
			})
			fmt.Print(err)
			fmt.Println(msgID)

			return c.JSON(http.StatusAccepted, make(map[string]string))
		}
	})
	app.Run(standard.New(":3000"))
}