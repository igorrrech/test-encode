package app

type Person struct {
	Id        int64  `json:"id" db:"id"`
	Email     string `json:"email" db:"email"`
	Phone     string `json:"phone" db:"phone"`
	FirstName string `json:"first-name" db:"first-name"`
	LastName  string `json:"last-name" db:"last-name"`
}
