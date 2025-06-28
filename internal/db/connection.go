package db

import (
	"context"
	"log"
	"os"

	"vinyl-store/ent" // путь зависит от твоего `go.mod`

	_ "github.com/lib/pq" // PostgreSQL драйвер
)

func Connect() *ent.Client {
	// Получаем строку подключения из переменной окружения
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set")
	}

	// Подключаемся к базе
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	// Выполняем миграции (создание таблиц)
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
