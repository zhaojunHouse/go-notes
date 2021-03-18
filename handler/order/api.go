package order

import (
	"github.com/labstack/echo"
)

type OrderHandlerInterface interface {
	GetOrder(ctx echo.Context) error
}

type Order struct {
}

func NewOrderHandler() OrderHandlerInterface {
	return &Order{}
}
