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
	insertUser        = `INSERT INTO users(id, first_name, last_name) VALUES ($1, $2, $3);`
	upsertHobby       = `INSERT INTO hobbies(name) VALUES (LOWER($1)) ON CONFLICT DO NOTHING;`
	insertUserHobbies = `INSERT INTO users_hobbies (user_id, hobby_id) VALUES($1 , (SELECT id FROM hobbies WHERE name=LOWER($2)))`
	getUserHobbies    = `SELECT name FROM users_hobbies LEFT JOIN hobbies ON hobbies.id=users_hobbies.hobby_id WHERE user_id=$1;`
	delUserHobbies    = `DELETE FROM users_hobbies where user_id=$1 and hobby_id=$2;`
)

func NewRepo() (*Repo, error) {
	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return &Repo{conn: conn}, nil
}

func (r *Repo) InsertUser(ctx context.Context, userID, firstName, lastName string) error {
	_, err := r.conn.Exec(ctx, insertUser, userID, firstName, lastName)
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

func (r *Repo) InsertUserHobbies(ctx context.Context, userID string, hobbies []string) error {
	var gp errgroup.Group

	for _, v := range hobbies {
		hobby := v
		gp.Go(func() error {
			_, err := r.conn.Exec(ctx, insertUserHobbies, userID, hobby)
			return err
		})
	}

	return gp.Wait()
}
