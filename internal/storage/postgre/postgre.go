package postgre

import (
	"context"
	"database/sql"

	"github.com/BelyaevEI/platform_common/pkg/closer"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func New(ctx context.Context, dsn string) (*sql.DB, error) {
	pg, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = pg.Ping()
	if err != nil {
		return nil, err
	}

	closer.Add(pg.Close)

	return pg, nil
}
