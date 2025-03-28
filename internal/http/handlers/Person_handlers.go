package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"test/app"
	"test/internal/http/middleware"
	l "test/internal/logic"

	"github.com/gocraft/dbr"
	"github.com/labstack/echo/v4"
)

func GetPersonList(
	personsGetter l.UseCaseExecuteReturnPersonList,
) echo.HandlerFunc {
	type (
		Response struct {
			Persons []app.Person `json:"persons"`
		}
	)
	return func(c echo.Context) error {
		var limit uint64
		if parsed, err := strconv.ParseUint(c.QueryParam(queryParamLimit), 10, 64); err == nil {
			limit = parsed
		}
		var offset uint64
		if parsed, err := strconv.ParseUint(c.QueryParam(queryParamOffset), 10, 64); err == nil {
			offset = parsed
		}
		searchString := c.QueryParam(queryParamSearch)
		var (
			session *dbr.Session
			ok      bool
		)
		if session, ok = c.Get(middleware.CSession).(*dbr.Session); !ok || session == nil {
			return ErrHaveNotSession
		}
		persons, err := personsGetter.Execute(session, limit, offset, searchString)
		if err != nil {
			return &echo.HTTPError{
				Code:     http.StatusInternalServerError,
				Message:  err.Error(),
				Internal: err,
			}
		}
		return c.JSON(http.StatusOK, Response{
			Persons: persons,
		})
	}
}
func GetPersonById(
	personGetter l.UseCaseExecuteReturnPerson,
) echo.HandlerFunc {
	type (
		Response struct {
			Person app.Person `json:"person"`
		}
	)
	return func(c echo.Context) error {
		var id uint64
		if parsed, err := strconv.ParseUint(c.Param(paramId), 10, 64); err == nil {
			id = parsed
		} else {
			return echo.ErrNotFound
		}
		var (
			session *dbr.Session
			ok      bool
		)
		if session, ok = c.Get(middleware.CSession).(*dbr.Session); !ok || session == nil {
			return ErrHaveNotSession
		}
		person, err := personGetter.Execute(session, id)
		if err != nil {
			return &echo.HTTPError{
				Code:     http.StatusInternalServerError,
				Message:  err.Error(),
				Internal: err,
			}
		}
		return c.JSON(http.StatusOK, Response{
			Person: person,
		})
	}
}
func CreatePerson(
	creater l.UseCaseExecutesPerson,
) echo.HandlerFunc {
	type (
		Request struct {
			Person app.Person `json:"person"`
		}
	)
	return func(c echo.Context) error {
		var (
			session *dbr.Session
			ok      bool
		)
		if session, ok = c.Get(middleware.CSession).(*dbr.Session); !ok || session == nil {
			return ErrHaveNotSession
		}
		var req Request
		if err := c.Bind(&req); err != nil || req.Person.IsEmpty() { //nothing to create if empty
			return c.NoContent(http.StatusBadRequest)
		}
		err := creater.Execute(session, req.Person)
		if err != nil {
			return &echo.HTTPError{
				Code:     http.StatusInternalServerError,
				Message:  err.Error(),
				Internal: err,
			}
		}
		return c.NoContent(http.StatusOK)
	}
}
func UpdatePerson(
	updater l.UseCaseExecutesPerson,
) echo.HandlerFunc {
	type (
		Request struct {
			Person app.Person `json:"person"`
		}
	)
	return func(c echo.Context) error {
		var id uint64
		if parsed, err := strconv.ParseUint(c.Param(paramId), 10, 64); err == nil {
			id = parsed
		} else {
			return echo.ErrNotFound
		}
		var (
			session *dbr.Session
			ok      bool
		)
		if session, ok = c.Get(middleware.CSession).(*dbr.Session); !ok || session == nil {
			return ErrHaveNotSession
		}
		var req Request
		if err := c.Bind(&req); err != nil || req.Person.IsEmpty() { //nothing to update if empty
			return c.NoContent(http.StatusBadRequest)
		}
		req.Person.Id = id
		err := updater.Execute(session, req.Person)
		if err != nil {
			return &echo.HTTPError{
				Code:     http.StatusInternalServerError,
				Message:  err.Error(),
				Internal: err,
			}
		}
		return c.NoContent(http.StatusOK)
	}
}
func DeletePerson(
	deleter l.UseCaseExecute,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		var id uint64
		if parsed, err := strconv.ParseUint(c.Param(paramId), 10, 64); err == nil {
			id = parsed
		} else {
			return echo.ErrNotFound
		}
		var (
			session *dbr.Session
			ok      bool
		)
		if session, ok = c.Get(middleware.CSession).(*dbr.Session); !ok || session == nil {
			return ErrHaveNotSession
		}
		err := deleter.Execute(session, id)
		if err != nil {
			return &echo.HTTPError{
				Code:     http.StatusInternalServerError,
				Message:  err.Error(),
				Internal: err,
			}
		}
		return c.NoContent(http.StatusOK)
	}
}

var (
	ErrHaveNotSession = errors.New("have not session in context")
)

const (
	paramId          = "id"
	queryParamLimit  = "limit"
	queryParamOffset = "offset"
	queryParamSearch = "search"
)
