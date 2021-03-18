package order

import (
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (o *Order) GetOrder(c echo.Context) error {
	log.Info("GetOrder method")
	return c.String(http.StatusOK, "GetOrder method")
}
