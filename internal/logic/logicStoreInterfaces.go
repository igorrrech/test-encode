package logic

import "test/app"

type PersonCreater interface {
	CreatePerson(person app.Person) error
}
type PersonGetter interface {
	GetPersonById(id uint64) (app.Person, error)
}
type PersonListGetter interface {
	GetPersons() ([]app.Person, error)
}
type PersonDeleter interface {
	DeletePerson(id uint64) error
}
type PersonUpdater interface {
	UpdatePerson(person app.Person) error
}
