package logic

import (
	"test/app"

	"github.com/gocraft/dbr"
)

type UseCaseExecute interface {
	Execute(session *dbr.Session, id uint64) error
}
type UseCaseExecutesPerson interface {
	Execute(session *dbr.Session, person app.Person) error
}
type UseCaseExecuteReturnPerson interface {
	Execute(session *dbr.Session, id uint64) (app.Person, error)
}
type UseCaseExecuteReturnPersonList interface {
	Execute(session *dbr.Session, limit uint64, offset uint64, search string) ([]app.Person, error)
}
