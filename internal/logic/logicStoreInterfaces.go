package logic

import (
	"test/app"

	"github.com/gocraft/dbr"
)

type PersonCreater interface {
	CreatePerson(tx *dbr.Tx, person app.Person) error
}
type PersonGetter interface {
	GetPersonById(tx *dbr.Tx, id uint64) (app.Person, error)
}
type PersonListGetter interface {
	GetPersons(tx *dbr.Tx, limit uint64, offset uint64, search string) ([]app.Person, error)
}
type PersonDeleter interface {
	DeletePerson(tx *dbr.Tx, id uint64) error
}
type PersonUpdater interface {
	UpdatePerson(tx *dbr.Tx, person app.Person) error
}
type PersonRepoInterface interface {
	PersonCreater
	PersonGetter
	PersonListGetter
	PersonDeleter
	PersonUpdater
}
type PersonRepoMock struct {
	Persons      []app.Person
	Person       app.Person
	CreateError  error
	GetError     error
	GetListError error
	UpdateError  error
	DeleteError  error
}

func (m PersonRepoMock) CreatePerson(tx *dbr.Tx, person app.Person) error {
	return m.CreateError
}
func (m PersonRepoMock) GetPersonById(tx *dbr.Tx, id uint64) (app.Person, error) {
	return m.Person, m.GetError
}
func (m PersonRepoMock) GetPersons(tx *dbr.Tx, limit uint64, offset uint64, search string) ([]app.Person, error) {
	return m.Persons, m.GetListError
}
func (m PersonRepoMock) DeletePerson(tx *dbr.Tx, id uint64) error {
	return m.DeleteError
}
func (m PersonRepoMock) UpdatePerson(tx *dbr.Tx, person app.Person) error {
	return m.UpdateError
}
