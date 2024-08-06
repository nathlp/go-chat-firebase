package routers

import (
	"github.com/labstack/echo/v4"
	"go-chat-firebase/handler"
)

type Router struct {
	message handler.MessageHandlerInterface
}

func NewRouter(message handler.MessageHandlerInterface) *Router {
	return &Router{
		message: message,
	}
}

func EndPoints(router *Router) {
	e := echo.New()
	e.GET("/health_check", router.message.SendMessage)
	e.POST("/send", router.message.SendMessage)
	e.GET("/messages", getMessages)
	e.Logger.Fatal(e.Start(":8080"))
}
