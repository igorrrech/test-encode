package logic

import (
	"test/app"

	"github.com/gocraft/dbr"
)

type UseCaseDeletePerson struct {
	deleter PersonDeleter
}

func NewUseCaseDeletePerson(
	p *PersonDeleter,
) *UseCaseDeletePerson {
	return &UseCaseDeletePerson{
		deleter: *p,
	}
}
func (u UseCaseDeletePerson) Execute(session *dbr.Session, id uint64) error {
	//ctx
	//tx
	return nil
}

type UseCaseCreatePerson struct {
	creater PersonCreater
}

func NewUseCaseCreatePerson(
	p *PersonCreater,
) *UseCaseCreatePerson {
	return &UseCaseCreatePerson{
		creater: *p,
	}
}
func (u UseCaseCreatePerson) Execute(session *dbr.Session, person app.Person) error {
	//validate
	//ctx
	//tx
	return nil
}

type UseCaseUpdatePerson struct {
	updater PersonUpdater
}

func NewUseCaseUpdatePerson(
	p *PersonUpdater,
) *UseCaseUpdatePerson {
	return &UseCaseUpdatePerson{
		updater: *p,
	}
}
func (u UseCaseUpdatePerson) Execute(session *dbr.Session, person app.Person) error {
	//validate
	//ctx
	//tx
	return nil
}

type UseCaseGetPersonById struct {
	getter PersonGetter
}

func NewUseCaseGetPersonById(
	p *PersonGetter,
) *UseCaseGetPersonById {
	return &UseCaseGetPersonById{
		getter: *p,
	}
}
func (u UseCaseGetPersonById) Execute(session *dbr.Session, id uint64) (app.Person, error) {
	//ctx
	//tx
	return app.Person{}, nil
}

type UseCaseGetPersonsList struct {
	listGetter PersonListGetter
}

func NewUseCaseGetPersonsList(
	p *PersonListGetter,
) *UseCaseGetPersonsList {
	return &UseCaseGetPersonsList{
		listGetter: *p,
	}
}
func (u UseCaseGetPersonsList) Execute(
	session *dbr.Session,
	limit uint64,
	offset uint64,
	search string,
) ([]app.Person, error) {
	//ctx
	//tx
	return nil, nil
}
