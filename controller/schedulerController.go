package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"smsSender/httputil"
	"smsSender/worker"
)

// StartScheduler godoc
//
//	@Summary		Start worker
//	@Description	Start worker
//	@Tags			worker
//	@Accept			json
//	@Produce		json
//	@Success		200     {object}	httputil.HTTPSuccess
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		404		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Router			/worker/start [put]
func (c *Controller) Start(ctx *gin.Context) {
	worker.Start()
	httputil.NewSuccess(ctx, http.StatusAccepted, "Scheduler started!")
}

// StopScheduler godoc
//
//	@Summary		Stop worker
//	@Description	Stop worker
//	@Tags			worker
//	@Accept			json
//	@Produce		json
//	@Success		200     {object}	httputil.HTTPSuccess
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		404		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Router			/worker/stop [put]
func (c *Controller) Stop(ctx *gin.Context) {
	err := worker.Stop()
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	} else {
		httputil.NewSuccess(ctx, http.StatusAccepted, "Scheduler stopped!")
	}
}
