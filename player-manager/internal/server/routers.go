package server

import (
	"player-manager/internal/server/auth"
	"player-manager/internal/server/log"
	"player-manager/internal/server/player"

	"github.com/gin-gonic/gin"
)

const (
	apiV1Prefix   = "/api/v1"
	playersPrefix = apiV1Prefix + "/players"
	logsPrefix    = apiV1Prefix + "/logs"
)

func (s *Server) setupSwaggerRoute(r *gin.Engine) {
	r.Static(apiV1Prefix+"/swagger", "/resources/swagger-ui-dist")

}

func (s *Server) setupLogRoutes(r *gin.Engine) {
	ctrl := log.NewController(s.db)

	logsSecretAPI := r.Group(logsPrefix, auth.SimpleAuthorizationMiddleware)
	{
		logsSecretAPI.GET("/", ctrl.GetLogs)
	}
}

func (s *Server) setupPlayerRoutes(r *gin.Engine) {
	ctrl := player.NewController(s.db)

	playersOpenAPI := r.Group(playersPrefix)
	{
		playersOpenAPI.GET("/:username", ctrl.GetPlayer)
		playersOpenAPI.GET("/:username/items", ctrl.ListItems)
	}

	playersSecretAPI := r.Group(playersPrefix, auth.SimpleAuthorizationMiddleware)
	{
		playersSecretAPI.PUT("/:username/items/:item", ctrl.AddItem)
		playersSecretAPI.DELETE("/:username/items/:item", ctrl.DeleteItem)

		playersSecretAPI.POST("/:username/action/build", ctrl.Build)
	}
}
