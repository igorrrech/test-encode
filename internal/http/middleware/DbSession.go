package middleware

import (
	pdb "test/persondb"

	"github.com/labstack/echo/v4"
)

func NewDbSessionMiddleware(p *pdb.ConnectionProvider) echo.MiddlewareFunc {
	//deps
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			//middleware code
			//create session and bind into echo context
			c.Set(CSession, p.GetSession(nil))
			return next(c)
		}
	}
}

const (
	CSession = "Session"
)
