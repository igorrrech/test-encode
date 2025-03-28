package logic

import (
	"test/app"

	"github.com/gocraft/dbr"
)

// Use cases mocks
type ExecuteMock struct {
	Error error
}

func (m ExecuteMock) Execute(s *dbr.Session, id uint64) error {
	return m.Error
}

type ExecutesPersonMock struct {
	Error error
}

func (m ExecutesPersonMock) Execute(s *dbr.Session, person app.Person) error {
	return m.Error
}

type ExecuteReturnPersonMock struct {
	Error  error
	Person app.Person
}

func (m ExecuteReturnPersonMock) Execute(s *dbr.Session, id uint64) (app.Person, error) {
	return m.Person, m.Error
}

type ExecuteReturnPersonListMock struct {
	Error   error
	Persons []app.Person
}

func (m ExecuteReturnPersonListMock) Execute(s *dbr.Session, limit uint64, offset uint64, search string) ([]app.Person, error) {
	return m.Persons, m.Error
}
