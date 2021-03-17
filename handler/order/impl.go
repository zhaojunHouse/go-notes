package order

import (
	"context"
	log "github.com/sirupsen/logrus"
)

func (o *Order) GetOrder(ctx context.Context, userID int64) interface{} {
	log.Info("GetOrder method")
	panic("implement me")
}

func (o *Order) UpdateOrder(ctx context.Context, userID int64) interface{} {
	log.Info("UpdateOrder Method")
	panic("implement me")
}
