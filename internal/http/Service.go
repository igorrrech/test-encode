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

	h "test/internal/http/handlers"
	m "test/internal/http/middleware"
	"test/internal/logic"
)

type Service struct {
	host string
	port string
	log  *logrus.Logger
	pr   logic.PersonRepoInterface
	sp   m.SessionProvider
}

func NewService(
	host string,
	port string,
	log *logrus.Logger,
	personRepo logic.PersonRepoInterface,
	sessionProvider m.SessionProvider,

) *Service {
	return &Service{
		host: host,
		port: port,
		log:  log,
		pr:   personRepo,
		sp:   sessionProvider,
	}
}
func (s Service) Run(ctx context.Context) {
	//use cases init
	creator := logic.NewUseCaseCreatePerson(s.pr)
	updater := logic.NewUseCaseUpdatePerson(s.pr)
	deleter := logic.NewUseCaseDeletePerson(s.pr)
	getter := logic.NewUseCaseGetPersonById(s.pr)

	listGetter := logic.NewUseCaseGetPersonsList(s.pr)
	e := echo.New()
	//logrus err handler adapter
	e.HTTPErrorHandler = m.NewLogrusErrorHandler(s.log)
	//logrus logging middleware adapter
	e.Use(m.NewLogrusMiddleware(s.log), middleware.Recover())

	g := e.Group("/person", m.NewDbSessionMiddleware(s.sp))
	g.GET("/", h.GetPersonList(listGetter))
	g.GET("/:id", h.GetPersonById(getter))
	g.POST("/", h.CreatePerson(creator))
	g.PUT("/:id", h.UpdatePerson(updater))
	g.DELETE("/:id", h.DeletePerson(deleter))

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
