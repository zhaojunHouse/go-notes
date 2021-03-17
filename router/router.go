package router

import (
	"github.com/labstack/echo"
	"go-notes/handler/order"
	"go-notes/handler/user"
)

func Router(
	userHandlerInterface user.UserHandlerInterface,
	orderHandlerInterface order.OrderHandlerInterface,
) *echo.Echo {
	engine := echo.New()
	engine.GET("/user", userHandlerInterface.GetUser)
	engine.GET("/order", orderHandlerInterface.GetOrder)
	return engine
}
