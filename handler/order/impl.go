package order

import (
	"context"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (o *Order) GetOrder(c echo.Context) error {
	log.Info("GetOrder method")
	return c.String(http.StatusOK, "GetOrder method")
}

func (o *Order) UpdateOrder(ctx context.Context, userID int64) interface{} {
	log.Info("UpdateOrder Method")
	panic("implement me")
}
