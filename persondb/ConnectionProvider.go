package persondb

import (
	"fmt"

	"github.com/gocraft/dbr"
	"github.com/sirupsen/logrus"
)

type ConnectionProvider struct {
	driver string
	dsn    string
	event  dbr.EventReceiver
	log    *logrus.Logger
	conn   *dbr.Connection
}

func NewConnectionProvider(
	driver string,
	dsn string,
	event dbr.EventReceiver,
	log *logrus.Logger,
) *ConnectionProvider {
	cp := &ConnectionProvider{
		driver: driver,
		dsn:    dsn,
		log:    log,
	}
	cp.mustCreateConnection()
	return cp
}
func (p *ConnectionProvider) mustCreateConnection() {
	conn, err := dbr.Open(p.driver, p.dsn, p.event)
	if err != nil {
		p.log.Fatal(fmt.Sprintf("connection open error:%s", err.Error()))
	}
	p.conn = conn
}
func (p *ConnectionProvider) GetSession(event dbr.EventReceiver) *dbr.Session {
	return p.conn.NewSession(event)
}
