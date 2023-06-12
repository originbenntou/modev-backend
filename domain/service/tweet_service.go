package service

import (
	"context"
	"fmt"
	"github.com/originbenntou/modev-backend/domain/entity"
	"github.com/originbenntou/modev-backend/domain/model"
	"github.com/originbenntou/modev-backend/domain/repository"
	vo "github.com/originbenntou/modev-backend/domain/value_object"
	"time"
)

type TweetService interface {
	FindByCategory(ctx context.Context, category *vo.Category) ([]*entity.TweetEntity, error)
}

type tweetService struct {
	repository.TweetRepository
}

func NewTweetService(t repository.TweetRepository) TweetService {
	return &tweetService{
		t,
	}
}

func (s *tweetService) FindByCategory(ctx context.Context, category *vo.Category) ([]*entity.TweetEntity, error) {
	tweetModels, err := s.TweetRepository.FindByCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	tweetEntities := make([]*entity.TweetEntity, len(tweetModels))
	for i, m := range tweetModels {
		tweetEntities[i] = modelToEntity(m)
	}

	return tweetEntities, nil
}

func modelToEntity(m *model.TweetModel) *entity.TweetEntity {
	category, err := vo.NewCategory(m.Category)
	if err != nil {
		fmt.Println("modelToEntity category")
		return nil
	}
	addDate, err := vo.NewAddDate(m.AddDate)
	if err != nil {
		fmt.Println("modelToEntity addDate")
		return nil
	}

	url, err := vo.NewURL(m.Url)
	if err != nil {
		fmt.Println("modelToEntity url")
		return nil
	}

	createdAt, err := time.Parse("2006-01-02T15:04:05Z", m.CreatedAt)
	if err != nil {
		fmt.Println("modelToEntity createdAt")
		return nil
	}

	updatedAt, err := time.Parse("2006-01-02T15:04:05Z", m.UpdatedAt)
	if err != nil {
		fmt.Println("modelToEntity updatedAt")
		return nil
	}

	// FIXME: 色々キモい
	// ってかValueにする必要あるのか？
	loc, _ := time.LoadLocation("Asia/Tokyo")

	return &entity.TweetEntity{
		Id:        m.Id,
		Category:  *category,
		AddDate:   *addDate,
		Url:       *url,
		Tags:      nil,
		CreatedAt: createdAt.In(loc),
		UpdatedAt: updatedAt.In(loc),
	}
}
