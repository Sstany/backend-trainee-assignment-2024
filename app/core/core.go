package core

import (
	"banney/app/db"
	"banney/app/handlers/banner"
	"banney/app/router"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	host     string
	dbClient *db.Client
	logger   *zap.Logger
}

func NewServer(host string, dbClient *db.Client, logger *zap.Logger) *Server {
	return &Server{
		host:     host,
		dbClient: dbClient,
		logger:   logger,
	}

}

func (r *Server) Start() {
	api := r.newAPI()
	api.Run(r.host)
}

func (r *Server) newAPI() *gin.Engine {
	engine := gin.New()

	pr := router.New(r.dbClient, r.logger.Named("router"))
	banner.AttachToGroup(pr, engine.Group("banner"))

	return engine
}
