package business

import (
	"context"

	"github.com/frisk038/livechat/business/models"
)

type repo interface {
	InsertUser(ctx context.Context, username, password string) error
	InsertHobbies(ctx context.Context, hobbies []string) error
	InsertUserHobbies(ctx context.Context, username string, hobbies []string) error
}

type BusinessProfile struct {
	repo repo
}

func NewBusinessProfile(repo repo) BusinessProfile {
	return BusinessProfile{repo: repo}
}

func (bp *BusinessProfile) CreateUser(ctx context.Context, user models.User) error {
	err := bp.repo.InsertUser(context.Background(), user.Name, user.Password)
	if err != nil {
		return err
	}

	err = bp.repo.InsertHobbies(context.Background(), user.Hobbies)
	if err != nil {
		return err
	}

	err = bp.repo.InsertUserHobbies(context.Background(), user.Name, user.Hobbies)
	if err != nil {
		return err
	}

	return nil
}
