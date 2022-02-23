package postgresqldao

import (
	"context"
	"database/sql"
	"log"
	"os"
	"sync"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PostgreSql struct {
	client *sql.DB
	ctx    context.Context
}

var once sync.Once
var instance *PostgreSql

func NewPostgreSql(ctx context.Context) *PostgreSql {

	once.Do(func() {
		instance = new(PostgreSql)
		instance.ctx = ctx
		instance.connect()
	})

	return instance
}

func (m *PostgreSql) connect() {
	var err error
	pgPass := os.Getenv("PGP")
	connectionstring := "host=130.61.54.93 port=49153 user=postgres dbname=paygateway sslmode=disable password=" + pgPass
	m.client, err = sql.Open("postgres", connectionstring)
	if err != nil {
		log.Println("Error connecting to database", err)
	}
}

func (m *PostgreSql) Disconnect() {
	m.client.Close()
}
