//	Player manager Api:
//	  version: 1.0.0
//	  title: Player Manager API
//	Schemes: http
//	Host: localhost:8080
//	BasePath: /api/v1
//	Produces:
//	  - application/json
//
// swagger:meta
package server

import (
	"database/sql"
	"fmt"
	"log"
	"player-manager/internal/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg config.Config
	db  *sql.DB
}

func New(cfg config.Config, db *sql.DB) *Server {
	return &Server{cfg, db}
}

func (s *Server) Start() error {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Recovery())

	s.setupSwaggerRoute(router)
	s.setupLogRoutes(router)
	s.setupPlayerRoutes(router)

	log.Println("Starting the server with the configuration:")
	log.Println(s.cfg.String())

	return router.Run(fmt.Sprintf(":%d", s.cfg.Port))
}
