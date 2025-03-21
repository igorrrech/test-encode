package logic

import (
	"test/app"

	"github.com/gocraft/dbr"
)

type UseCaseDeletePerson struct {
}

func NewUseCaseDeletePerson(
	p *PersonDeleter,
) *UseCaseDeletePerson {
	return nil
}
func (u UseCaseDeletePerson) Execute(session *dbr.Session, id uint64) error {
	//ctx
	//tx
	return nil
}

type UseCaseCreatePerson struct {
}

func NewUseCaseCreatePerson(
	p *PersonCreater,
) *UseCaseCreatePerson {
	return nil
}
func (u UseCaseCreatePerson) Execute(session *dbr.Session, person app.Person) error {
	//validate
	//ctx
	//tx
	return nil
}

type UseCaseUpdatePerson struct {
}

func NewUseCaseUpdatePerson(
	p *PersonUpdater,
) *UseCaseUpdatePerson {
	return nil
}
func (u UseCaseUpdatePerson) Execute(session *dbr.Session, person app.Person) error {
	//validate
	//ctx
	//tx
	return nil
}

type UseCaseGetPersonById struct {
}

func NewUseCaseGetPersonById(
	p *PersonGetter,
) *UseCaseGetPersonById {
	return nil
}

type UseCaseGetPersonsList struct {
}

func NewUseCaseGetPersonsList(
	p *PersonListGetter,
) *UseCaseGetPersonsList {
	return nil
}
