package logic

import (
	"errors"
	"regexp"
	"test/app"
	db "test/persondb"

	"github.com/gocraft/dbr"
)

type UseCaseDeletePerson struct {
	deleter db.PersonDeleter
}

func NewUseCaseDeletePerson(
	p db.PersonDeleter,
) *UseCaseDeletePerson {
	return &UseCaseDeletePerson{
		deleter: p,
	}
}
func (u UseCaseDeletePerson) Execute(session *dbr.Session, id uint64) error {
	//ctx
	//tx
	tx, err := session.Begin()
	defer tx.RollbackUnlessCommitted()
	if err != nil {
		//err handle
		return errors.Join(ErrTxCreation, err)
	}
	err = u.deleter.DeletePerson(tx, id)
	if err != nil {
		//err handle
		return errors.Join(ErrExecute, err)
	}
	if err = tx.Commit(); err != nil {
		//err handle
		return errors.Join(ErrTxCommit, err)
	}
	return nil
}

type UseCaseCreatePerson struct {
	creater db.PersonCreater
}

func NewUseCaseCreatePerson(
	p db.PersonCreater,
) *UseCaseCreatePerson {
	return &UseCaseCreatePerson{
		creater: p,
	}
}
func (u UseCaseCreatePerson) Execute(session *dbr.Session, person app.Person) error {
	//validate
	if err := person.Validate(); err != nil {
		//err handle
		return errors.Join(ErrValidatePerson, err)
	}
	//ctx
	//tx
	tx, err := session.Begin()
	defer tx.RollbackUnlessCommitted()
	if err != nil {
		//err handle
		return errors.Join(ErrTxCreation, err)
	}
	person.Print()
	err = u.creater.CreatePerson(tx, person)
	if err != nil {
		//err handle
		return errors.Join(ErrExecute, err)
	}
	if err = tx.Commit(); err != nil {
		//err handle
		return errors.Join(ErrTxCommit, err)
	}
	return nil
}

type UseCaseUpdatePerson struct {
	updater db.PersonUpdater
}

func NewUseCaseUpdatePerson(
	p db.PersonUpdater,
) *UseCaseUpdatePerson {
	return &UseCaseUpdatePerson{
		updater: p,
	}
}
func (u UseCaseUpdatePerson) Execute(session *dbr.Session, person app.Person) error {
	//validate
	if err := person.Validate(); err != nil {
		//err handle
		return errors.Join(ErrValidatePerson, err)
	}
	//ctx
	//tx
	tx, err := session.Begin()
	defer tx.RollbackUnlessCommitted()
	if err != nil {
		//err handle
		return errors.Join(ErrTxCreation, err)
	}
	err = u.updater.UpdatePerson(tx, person)
	if err != nil {
		//err handle
		return errors.Join(ErrExecute, err)
	}
	if err = tx.Commit(); err != nil {
		//err handle
		return errors.Join(ErrTxCommit, err)
	}
	return nil
}

type UseCaseGetPersonById struct {
	getter db.PersonGetter
}

func NewUseCaseGetPersonById(
	p db.PersonGetter,
) *UseCaseGetPersonById {
	return &UseCaseGetPersonById{
		getter: p,
	}
}
func (u UseCaseGetPersonById) Execute(session *dbr.Session, id uint64) (app.Person, error) {
	var res app.Person
	//ctx
	//tx
	tx, err := session.Begin()
	defer tx.RollbackUnlessCommitted()
	if err != nil {
		//err handle
		return res, errors.Join(ErrTxCreation, err)
	}
	res, err = u.getter.GetPersonById(tx, id)
	if err != nil {
		//err handle
		return res, errors.Join(ErrExecute, err)
	}
	if err = tx.Commit(); err != nil {
		//err handle
		return res, errors.Join(ErrTxCommit, err)
	}
	return res, nil
}

type UseCaseGetPersonsList struct {
	listGetter db.PersonListGetter
}

func NewUseCaseGetPersonsList(
	p db.PersonListGetter,
) *UseCaseGetPersonsList {
	return &UseCaseGetPersonsList{
		listGetter: p,
	}
}
func (u UseCaseGetPersonsList) Execute(
	session *dbr.Session,
	limit uint64,
	offset uint64,
	search string,
) ([]app.Person, error) {
	var res []app.Person
	//validate search
	if !validateSerach(search) {
		return res, ErrValidateSearch
	}
	//ctx
	//tx
	tx, err := session.Begin()
	defer tx.RollbackUnlessCommitted()
	if err != nil {
		//err handle
		return res, errors.Join(ErrTxCreation, err)
	}
	res, err = u.listGetter.GetPersons(tx, limit, offset, search)
	if err != nil {
		//err handle
		return res, errors.Join(ErrExecute, err)
	}
	if err = tx.Commit(); err != nil {
		//err handle
		return res, errors.Join(ErrTxCommit, err)
	}
	return res, nil
}
func validateSerach(search string) bool {
	//regex
	//empty string is also valid
	sqlInj := regexp.MustCompile(`(\s*([\0\'\"\n\r\t\%\_\\]*\s*(((select\s*.+\s*from\s*.+)|(insert\s*.+\s*into\s*.+)|(update\s*.+\s*set\s*.+)|(delete\s*.+\s*from\s*.+)|(drop\s*.+)|(truncate\s*.+)|(alter\s*.+)|(exec\s*.+)|(\s*(all|any|not|and|between|in|like|or|some|contains|containsall|containskey)\s*.+[\=\>\<=\!\~]+.+)|(let\s+.+[\=]\s*.*)|(begin\s*.*\s*end)|(\s*[\/\*]+\s*.*\s*[\*\/]+)|(\s*(\-\-)\s*.*\s+)|(\s*(contains|containsall|containskey)\s+.*)))(\s*[\;]\s*)*)+)`)
	return search == "" || !sqlInj.MatchString(search)
}

var (
	ErrTxCreation     = errors.New("cant create transaction")
	ErrTxCommit       = errors.New("cant commit transaction")
	ErrValidateSearch = errors.New("search is not valid")
	ErrValidatePerson = errors.New("person is not valid")
	ErrExecute        = errors.New("repo cant commit operation")
)
