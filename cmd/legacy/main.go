package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"entgo.io/bug/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/go-sql-driver/mysql"
)

func main() {
	config := mysql.NewConfig()
	config.User = "root"
	config.Passwd = "pass"
	config.DBName = "test"
	config.Net = "tcp"
	config.Addr = fmt.Sprintf("%s:%s", "127.0.0.1", "3308")

	client, err := ent.Open(dialect.MySQL, config.FormatDSN())
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	opts := []schema.MigrateOption{
		schema.WithAtlas(false),
	}
	if err := client.Schema.WriteTo(context.Background(), os.Stdout, opts...); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}
	if err := client.Schema.Create(context.Background(), opts...); err != nil {
		log.Fatalf("failed createing schema resources: %v", err)
	}
}
