package middleware

import (
	"github.com/gocraft/dbr"
	"github.com/labstack/echo/v4"
)

type SessionProvider interface {
	GetSession(event dbr.EventReceiver) *dbr.Session
}
type SessionProviderMock struct {
	Session *dbr.Session
}

func (m SessionProviderMock) GetSession(event dbr.EventReceiver) *dbr.Session {
	return m.Session
}
func NewDbSessionMiddleware(p SessionProvider) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			//middleware code
			//create session and bind into echo context
			session := p.GetSession(nil)
			c.Set(CSession, session)
			return next(c)
		}
	}
}

const (
	CSession = "Session"
)
