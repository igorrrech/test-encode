package logic

import "github.com/gocraft/dbr"

type UseCase interface {
	execute(dbr.Session)
}
