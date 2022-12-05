package main

import (
	"fmt"
	"github.com/MuhammadyusufAdhamov/booking/api"
	"github.com/MuhammadyusufAdhamov/booking/storage"
	"github.com/go-redis/redis/v9"
	"log"

	"github.com/MuhammadyusufAdhamov/booking/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load(".")

	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)
	fmt.Println("-------------------", psqlUrl)
	psqlConn, err := sqlx.Connect("postgres", psqlUrl)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Addr,
	})

	strg := storage.NewStoragePg(psqlConn)
	inMemory := storage.NewInMemoryStorage(rdb)

	apiServer := api.New(&api.RouterOptions{
		Cfg:      &cfg,
		Storage:  strg,
		InMemory: inMemory,
	})

	err = apiServer.Run(cfg.HttpPort)
	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

	log.Print("Server stopped")
}
