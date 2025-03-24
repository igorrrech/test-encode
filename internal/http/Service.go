package http

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"

	"test/app"
	h "test/internal/http/handlers"
	m "test/internal/http/middleware"
	l "test/internal/logic"
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
	return &Service{
		host: host,
		port: port,
		log:  log,
	}
}
func (s Service) Run(ctx context.Context) {
	e := echo.New()
	//logrus err handler adapter
	e.HTTPErrorHandler = m.NewLogrusErrorHandler(s.log)
	//logrus logging middleware adapter
	e.Use(m.NewLogrusMiddleware(s.log), middleware.Recover())

	g := e.Group("/person", m.NewDbSessionMiddleware(m.SessionProviderMock{Session: nil}))
	g.GET("/", h.GetPersonList(l.ExecuteReturnPersonListMock{
		Error: nil,
		Persons: []app.Person{
			{Id: 1},
			{Id: 2},
			{Id: 3},
		},
	}))
	g.GET("/:id", h.GetPersonById(l.ExecuteReturnPersonMock{
		Error: nil,
		Person: app.Person{
			Id: 1,
		},
	}))
	g.POST("/", h.CreatePerson(l.ExecutesPersonMock{
		Error: nil,
	}))
	g.PUT("/:id", h.UpdatePerson(l.ExecutesPersonMock{
		Error: nil,
	}))
	g.DELETE("/:id", h.DeletePerson(l.ExecuteMock{
		Error: nil,
	}))

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()
	// Start server
	hostport := net.JoinHostPort(s.host, s.port)
	s.log.Info(fmt.Sprintf("starting server at %s ...", hostport))
	go func() {
		if err := e.Start(hostport); err != nil && err != http.ErrServerClosed {
			s.log.Fatal("shutting down the server")
		}
	}()
	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		s.log.Fatal(err.Error())
		return
	}
	s.log.Info("server is stopped")
}
