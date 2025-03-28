package persondb

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
