package playlist

import (
	"context"
	"database/sql"
	"log"

	"github.com/BelyaevEI/platform_common/pkg/closer"
	"github.com/BelyaevEI/playlist/internal/model"
	"github.com/BelyaevEI/playlist/internal/storage/postgre"
)

// PlaylistRepository represents a repository for playlist entities.
type PlaylistRepository interface {
	GetFirstSongOfUser(ctx context.Context, login string) (model.Song, error)
	GetNextSongOfUser(ctx context.Context, login string, curID, nextID int64) (model.Song, error)
}

type repo struct {
	db *sql.DB
}

func NewRepository(ctx context.Context, dsn string) PlaylistRepository {
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
