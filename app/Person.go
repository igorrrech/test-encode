package app

import (
	"errors"
	"fmt"
	"regexp"
)

type Person struct {
	Id        uint64 `json:"id" db:"id"`
	Email     string `json:"email" db:"email"`
	Phone     string `json:"phone" db:"phone"`
	FirstName string `json:"first-name" db:"first_name"`
	LastName  string `json:"last-name" db:"last_name"`
}

func (p Person) Validate() error {
	if !validateEmail(p.Email) {
		return ErrEmaiValidate
	}
	if !validatePhone(p.Phone) {
		return ErrPhoneValidate
	}
	if !validateName(p.FirstName) {
		return ErrFirstNameValidate
	}
	if !validateName(p.LastName) {
		return ErrLastNameValidate
	}
	return nil
}
func validateEmail(email string) bool {
	validEmail := regexp.MustCompile(`([a-zA-Z0-9._-]+@[a-zA-Z0-9._-]+\.[a-zA-Z0-9_-]+)$`)
	return email == "" || validEmail.MatchString(email)
}
func validatePhone(phone string) bool {
	validPhone := regexp.MustCompile(`^[+]?[(]{0,1}[0-9]{1,4}[)]{0,1}[\s\./0-9]*$`)
	return phone == "" || validPhone.MatchString(phone)
}
func validateName(name string) bool {
	//try to detect sqlinj
	sqlInj := regexp.MustCompile(`(\s*([\0\'\"\n\r\t\%\_\\]*\s*(((select\s*.+\s*from\s*.+)|(insert\s*.+\s*into\s*.+)|(update\s*.+\s*set\s*.+)|(delete\s*.+\s*from\s*.+)|(drop\s*.+)|(truncate\s*.+)|(alter\s*.+)|(exec\s*.+)|(\s*(all|any|not|and|between|in|like|or|some|contains|containsall|containskey)\s*.+[\=\>\<=\!\~]+.+)|(let\s+.+[\=]\s*.*)|(begin\s*.*\s*end)|(\s*[\/\*]+\s*.*\s*[\*\/]+)|(\s*(\-\-)\s*.*\s+)|(\s*(contains|containsall|containskey)\s+.*)))(\s*[\;]\s*)*)+)`)
	return name == "" || !sqlInj.MatchString(name)
}
func (p Person) Print() {
	fmt.Printf("person\n")
	fmt.Println(p.Id)
	fmt.Println(p.Email)
	fmt.Println(p.Phone)
	fmt.Println(p.FirstName)
	fmt.Println(p.LastName)
}
func (p Person) IsEmpty() bool {
	return p.Email == "" && p.Phone == "" && p.FirstName == "" && p.LastName == ""
}

var (
	ErrEmaiValidate      = errors.New("email is not valid")
	ErrPhoneValidate     = errors.New("phone is not valid")
	ErrFirstNameValidate = errors.New("first name is not valid")
	ErrLastNameValidate  = errors.New("last name is not valid")
)
