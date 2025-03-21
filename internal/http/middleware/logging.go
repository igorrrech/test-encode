package middleware

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func NewLogrusMiddleware(log *logrus.Logger) echo.MiddlewareFunc {
	// skippedRouts := make(map[string]struct{})
	// skippedRouts["/health"] = struct{}{}

	// skipper := func(c echo.Context) bool {
	// 	_, ok := skippedRouts[c.Request().URL.Path]
	// 	return ok
	// }

	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		//Skipper:   skipper,
		BeforeNextFunc: func(c echo.Context) {
			c.Set("timeSince", time.Now())
		},
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			timeSince, ok := c.Get("timeSince").(time.Time)
			if !ok {
				timeSince = time.Time{}
			}
			log.WithFields(logrus.Fields{
				"URI":    v.URI,
				"status": v.Status,
				"time":   time.Since(timeSince),
			}).Info("request")
			return nil
		},
	})
}
func NewLogrusErrorHandler(log *logrus.Logger) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}
		echoErr, ok := err.(*echo.HTTPError)
		if !ok {
			echoErr = &echo.HTTPError{}
			echoErr.Code = http.StatusInternalServerError
			echoErr.Message = ""
			echoErr.Internal = err
		}
		log.WithFields(logrus.Fields{
			"code": echoErr.Code,
			"msg":  echoErr.Message,
			"err":  echoErr.Internal.Error(),
		}).Error("Error")

	}
}
