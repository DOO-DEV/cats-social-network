package db

import (
	"context"
	"database/sql"
	"meower/schema"
)

type PostgresRepository struct {
	db *sql.DB
}

func New(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return &PostgresRepository{db: db}, nil
}

func (d PostgresRepository) InsertMeow(ctx context.Context, meow schema.Meow) error {
	if _, err := d.db.ExecContext(ctx, `insert into meows values($1, $2, $3)`, meow.ID, meow.Body, meow.CreatedAt); err != nil {
		return err
	}

	return nil

}

func (d PostgresRepository) ListMeows(ctx context.Context, skip, take uint64) ([]schema.Meow, error) {
	rows, err := d.db.QueryContext(ctx, `select * from meows order by id desc offset $1 limit $2`, skip, take)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	meows := make([]schema.Meow, 0)
	for rows.Next() {
		meow := schema.Meow{}
		if err := rows.Scan(&meow.ID, &meow.Body, &meow.CreatedAt); err != nil {
			return nil, err
		}

		meows = append(meows, meow)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return meows, nil
}
