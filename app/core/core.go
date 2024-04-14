package core

import (
	"banney/app/db"
	"banney/app/handlers/auth"
	"banney/app/handlers/banner"
	"banney/app/router"
	"context"
	"errors"
	"net/http"
	"sync"

	userbanner "banney/app/handlers/user_banner"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	host     string
	dbClient *db.Client
	logger   *zap.Logger
	srv      *http.Server
	wg       *sync.WaitGroup
}

func NewServer(host string, dbClient *db.Client, logger *zap.Logger) *Server {
	return &Server{
		host:     host,
		dbClient: dbClient,
		logger:   logger,
		srv: &http.Server{
			Addr: host,
		},
		wg: &sync.WaitGroup{},
	}

}

func (r *Server) Start() {
	r.wg.Add(1)

	api := r.newAPI()

	r.srv.Handler = api.Handler()

	go func() {
		defer r.wg.Done()

		if err := r.srv.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				panic(err)
			}
		}
		r.logger.Info("server is stopped")
	}()
}

func (r *Server) Stop(ctx context.Context) error {
	if err := r.srv.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}

func (r *Server) Wait() {
	r.wg.Wait()
}
func (r *Server) newAPI() *gin.Engine {
	engine := gin.New()

	pr := router.New(r.dbClient, r.logger.Named("router"))
	banner.AttachToGroup(pr, engine.Group("banner"))
	auth.AttachToGroup(pr, engine.Group("auth"))
	userbanner.AttachToGroup(pr, engine.Group("user_banner"))

	return engine
}
