package auth

import (
	"context"
	"database/sql"
	"log"

	"github.com/BelyaevEI/platform_common/pkg/closer"
	"github.com/BelyaevEI/playlist/internal/storage/postgre"
)

// AuthRepository represents a repository for auth entities.
type AuthRepository interface {
	CreateUser(ctx context.Context, login, hashPassword, secretKey string) error
	CheckLoginUnique(context.Context, string) error
}

type repo struct {
	db *sql.DB
}

func NewRepository(ctx context.Context, dsn string) AuthRepository {
	pg, err := postgre.New(ctx, dsn)
	if err != nil {
		log.Fatalf("failed connet to database: %s", err.Error())
	}

	err = pg.Ping()
	if err != nil {
		log.Fatalf("ping error: %s", err.Error())
	}

	closer.Add(pg.Close)

	return &repo{db: pg}
}
