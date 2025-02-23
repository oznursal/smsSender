package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
	"net/http"
	"smsSender/db"
	"smsSender/model"
)

// ListMessages godoc
//
//	@Summary		List messages
//	@Description	get messages
//	@Tags			messages
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		model.Message
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/messages [get]
func (c *Controller) ListMessages(ctx *gin.Context) {
	messages, err := db.GetSentMessages()
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	if messages == nil {
		messages = []model.Message{}
	}
	ctx.JSON(http.StatusOK, messages)
}
