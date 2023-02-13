package log

import (
	"net/http"
	"player-manager/internal/server/errors"

	"github.com/gin-gonic/gin"
)

// swagger:route GET /logs logs logs
//
// Gets logs.
//
//	    Security:
//	      token:
//
//		Responses:
//		  200: LogsResponse
//		  400: CommonError
func (ctrl *Controller) GetLogs(c *gin.Context) {
	logs, err := ctrl.logRepository.GetLogs(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.NewError(http.StatusBadRequest, "failed to get logs", err))
		return
	}

	c.JSON(http.StatusOK, ToLogsResponse(logs))
}
