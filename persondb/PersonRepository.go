package persondb

import (
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
	//add limit, offset and search in 'Where like'
	//check can it loaded into person array
	_, err := tx.Select("*").From(r.tableName).Load(persons)
	return persons, err
}
func (r PersonRepository) GetPersonById(tx *dbr.Tx, id uint64) (app.Person, error) {
	var result app.Person
	_, err := tx.Select("*").From(r.tableName).Where("id = ?", id).Load(&result)
	return result, err
}
func (r PersonRepository) DeletePerson(tx *dbr.Tx, id uint64) error {
	_, err := tx.DeleteFrom(r.tableName).Where("id = ?", id).Exec()
	return err
}
func (r PersonRepository) UpdatePerson(tx *dbr.Tx, person app.Person) error {
	mapToUpdate := make(map[string]interface{})
	mapToUpdate["email"] = person.Email
	mapToUpdate["phone"] = person.Phone
	mapToUpdate["first_name"] = person.FirstName
	mapToUpdate["last_name"] = person.LastName
	_, err := tx.Update(r.tableName).Where("Id = ?", person.Id).SetMap(mapToUpdate).Exec()
	return err

}
func (r PersonRepository) CreatePerson(tx *dbr.Tx, person app.Person) error {
	_, err := tx.InsertInto(r.tableName).Record(&person).Exec()
	return err
}
