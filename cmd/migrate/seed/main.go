package main

import (
	"log"

	"github.com/The-yohanna/social/internal/db"
	"github.com/The-yohanna/social/internal/env"
	"github.com/The-yohanna/social/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgresql://admin:password@localhost:5432/social?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store, conn)
}
