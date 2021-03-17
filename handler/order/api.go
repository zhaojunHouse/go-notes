package order

import "context"

type OrderHandlerInterface interface {
	GetOrder(ctx context.Context, userID int64) interface{}
	UpdateOrder(ctx context.Context, userID int64) interface{}
}

type Order struct {
}

func NewUserHandler() OrderHandlerInterface {
	return &Order{}
}
