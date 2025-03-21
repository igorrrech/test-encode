package http

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	h "test/internal/http/handlers"
	m "test/internal/http/middleware"
)

type Service struct {
	host string
	port string
	log  *logrus.Logger
}

func NewService(
	host string,
	port string,
	log *logrus.Logger,

) *Service {
	return nil
}
func (s Service) Run(ctx context.Context) {
	e := echo.New()
	//logrus err handler adapter
	e.HTTPErrorHandler = m.NewLogrusErrorHandler(s.log)
	//logrus logging middleware adapter
	e.Use(m.NewLogrusMiddleware(s.log))
	// e.GET("/", func(c echo.Context) error {
	// 	time.Sleep(5 * time.Second)
	// 	return c.JSON(http.StatusOK, "OK")
	// })
	g := e.Group("/person") //, m.NewDbSessionMiddleware())
	g.GET("/", h.GetPersonList())
	g.GET("/:id", h.GetPersonById())
	g.POST("/", h.CreatePerson())
	g.PUT("/:id", h.UpdatePerson())
	g.DELETE("/:id", h.DeletePerson())

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()
	// Start server
	hostport := net.JoinHostPort(s.host, s.port)
	e.Logger.Info("starting server at %s ...", hostport)
	go func() {
		if err := e.Start(hostport); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()
	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
		return
	}
	e.Logger.Info("server is stopped")
}
