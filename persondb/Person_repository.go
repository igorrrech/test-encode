package persondb

import (
	"errors"
	"test/app"

	"github.com/gocraft/dbr"
)

type PersonRepository struct {
	tableName string
}

func NewPersonRepository(
	tableName string,
) *PersonRepository {
	return &PersonRepository{
		tableName: tableName,
	}

}
func (r PersonRepository) GetPersons(tx *dbr.Tx, limit uint64, offset uint64, search string) ([]app.Person, error) {
	var persons []app.Person
	stmt := tx.Select("*").From(r.tableName)
	//add limit, offset and search in 'Where like'
	if limit > 0 {
		stmt = stmt.Limit(limit)
	}
	if offset > 0 {
		stmt.Offset(offset)
	}
	if search != "" {
		stmt.Where("first_name LIKE '%' || ? || '%'", search)
	}
	//check can it loaded into person array
	_, err := stmt.Load(&persons)
	return persons, err
}
func (r PersonRepository) GetPersonById(tx *dbr.Tx, id uint64) (app.Person, error) {
	var result app.Person
	_, err := tx.Select("*").From(r.tableName).Where("id = ?", id).Load(&result)
	if result.Id == 0 {
		err = ErrNoSuchUser
	}
	return result, err
}
func (r PersonRepository) DeletePerson(tx *dbr.Tx, id uint64) error {
	var selected app.Person
	tx.Select("*").From(r.tableName).Where("id =?", id).Load(&selected)
	if selected.Id == 0 {
		return ErrNoSuchUser
	}
	_, err := tx.DeleteFrom(r.tableName).Where("id = ?", id).Exec()
	return err
}
func (r PersonRepository) UpdatePerson(tx *dbr.Tx, person app.Person) error {
	mapToUpdate := make(map[string]interface{})
	mapToUpdate["email"] = person.Email
	mapToUpdate["phone"] = person.Phone
	mapToUpdate["first_name"] = person.FirstName
	mapToUpdate["last_name"] = person.LastName
	var selected app.Person
	tx.Select("*").From(r.tableName).Where("id =?", person.Id).Load(&selected)
	if selected.Id == 0 {
		return ErrNoSuchUser
	}
	_, err := tx.Update(r.tableName).Where("id = ?", person.Id).SetMap(mapToUpdate).Exec()
	return err

}
func (r PersonRepository) CreatePerson(tx *dbr.Tx, person app.Person) error {
	_, err := tx.InsertInto(r.tableName).Columns("email", "phone", "first_name", "last_name").Record(&person).Exec()
	return err
}

var (
	ErrNoSuchUser = errors.New("no such user")
)
