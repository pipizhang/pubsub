package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/pipizhang/pubsub/pkg"
)

type (
	brokerController struct {
		baseController
	}

	pubRequest struct {
		Channel string `json:"channel" form:"channel"`
		Message string `json:"message" form:"message"`
	}
)

var (
	BrokerController = &brokerController{}
)

func (b brokerController) Publish(c echo.Context) (err error) {
	// Validation
	req := new(pubRequest)
	if err = c.Bind(req); err != nil {
		return b.throwError(http.StatusBadRequest, "Invalid parameters")
	}
	if req.Message == "" {
		return b.throwError(http.StatusUnprocessableEntity, "Invalid message")
	}

	ch, isFound := pkg.Conf.GetChannel(req.Channel)
	if !isFound {
		return b.throwError(http.StatusUnprocessableEntity, "Not found channel")
	}

	client := pkg.NewPSClient()
	defer client.Disconnect(1)

	if client.Connect() != nil {
		return b.throwError(http.StatusInternalServerError, "PubSub service is unavailable")
	}

	err = client.Publish(ch.Key, req.Channel, req.Message)

	return c.JSON(http.StatusOK, "success")
}
