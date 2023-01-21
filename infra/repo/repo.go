package repo

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/sync/errgroup"
)

type Repo struct {
	conn *pgxpool.Pool
}

const (
	insertUser        = `INSERT INTO users(name, password) VALUES ($1, $2);`
	upsertHobby       = `INSERT INTO hobbies(name) VALUES (LOWER($1)) ON CONFLICT DO NOTHING;`
	insertUserHobbies = `INSERT INTO users_hobbies (user_id, hobby_id) 
	SELECT (SELECT id FROM users WHERE name=$1), id 
	FROM hobbies 
	WHERE name=ANY($2);`
)

func NewRepo() (*Repo, error) {
	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return &Repo{conn: conn}, nil
}

func (r *Repo) InsertUser(ctx context.Context, username, password string) error {
	_, err := r.conn.Exec(ctx, insertUser, username, password)
	return err
}

func (r *Repo) InsertHobbies(ctx context.Context, hobbies []string) error {
	var gp errgroup.Group

	for _, v := range hobbies {
		hobby := v
		gp.Go(func() error {
			_, err := r.conn.Exec(ctx, upsertHobby, hobby)
			return err
		})
	}

	return gp.Wait()
}

func (r *Repo) InsertUserHobbies(ctx context.Context, username string, hobbies []string) error {
	_, err := r.conn.Exec(ctx, insertUserHobbies, username, hobbies)
	return err
}
