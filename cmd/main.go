package main

import (
	"context"
	"test/internal/config"
	"test/internal/http"
	"test/persondb"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

//	&logic.PersonRepoMock{
//		Persons: []app.Person{
//			{Id: 1, Phone: "+79411037894", Email: "igorek@mail.com", FirstName: "Igorek", LastName: "Igorkov"},
//			{Id: 2, Phone: "+79511036844", Email: "vlad@yandex.com", FirstName: "Vlad", LastName: "Geroin"},
//			{Id: 3, Phone: "+79511337864", Email: "aye228@gmail.com", FirstName: "Петя", LastName: "Кувалда"},
//			{Id: 4, Phone: "+79111037894", Email: "lena322@yandex.com", FirstName: "Лена", LastName: "Янач"},
//		},
//		Person:       app.Person{Id: 1, Phone: "+79411037894", Email: "igorek@mail.com", FirstName: "Igorek", LastName: "Igorkov"},
//		CreateError:  nil,
//		UpdateError:  nil,
//		GetError:     nil,
//		GetListError: nil,
//		DeleteError:  nil,
//	}
func main() {
	cfg := config.MustLoadConfig("./config.json")
	logger := logrus.New()

	sp := persondb.NewConnectionProvider(
		"postgres",
		cfg.Dsn,
		nil,
		logger,
	)

	pr := persondb.NewPersonRepository("persons")

	svc := http.NewService(
		cfg.Host,
		cfg.Port,
		logger,
		pr,
		sp,
	)

	ctx := context.Background()
	svc.Run(ctx)
}
