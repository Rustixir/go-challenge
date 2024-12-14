package provider

import (
	"context"
	"sync"

	"entgo.io/ent/dialect"
	"github.com/Rustixir/go-challenge/ent"
	"github.com/Rustixir/go-challenge/pkg/config"
	_ "github.com/mattn/go-sqlite3"
)

var entOnce sync.Once
var entClient *ent.Client

func GetSqliteEnt() *ent.Client {
	entOnce.Do(func() {
		client, err := ent.Open(dialect.SQLite, config.Config.DB.SQL.DSN)
		if err != nil {
			panic(err)
		}
		err = client.Schema.Create(context.Background())
		if err != nil {
			panic(err)
		}
		entClient = client
	})
	return entClient
}
