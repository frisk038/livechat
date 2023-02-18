package business

import (
	"context"

	"github.com/frisk038/livechat/business/models"
)

type repo interface {
	InsertUser(ctx context.Context, userID, firstName, lastName string) error
	InsertHobbies(ctx context.Context, hobbies []string) error
	InsertUserHobbies(ctx context.Context, userID string, hobbies []string) error
}

type BusinessProfile struct {
	repo repo
}

func NewBusinessProfile(repo repo) BusinessProfile {
	return BusinessProfile{repo: repo}
}

func (bp *BusinessProfile) CreateUser(ctx context.Context, user models.User) error {
	err := bp.repo.InsertUser(context.Background(), user.ID, user.FirsttName, user.LastName)
	if err != nil {
		return err
	}

	err = bp.repo.InsertHobbies(context.Background(), user.Hobbies)
	if err != nil {
		return err
	}

	err = bp.repo.InsertUserHobbies(context.Background(), user.ID, user.Hobbies)
	if err != nil {
		return err
	}

	return nil
}
