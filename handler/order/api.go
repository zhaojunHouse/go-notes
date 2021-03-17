package order

import (
	"context"
	"github.com/labstack/echo"
)

type OrderHandlerInterface interface {
	GetOrder(ctx echo.Context) error
	UpdateOrder(ctx context.Context, userID int64) interface{}
}

type Order struct {
}

func NewOrderHandler() OrderHandlerInterface {
	return &Order{}
}
