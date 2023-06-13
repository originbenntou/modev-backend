package service

import (
	"context"
	"github.com/originbenntou/modev-backend/domain/entity"
	"github.com/originbenntou/modev-backend/domain/model"
	"github.com/originbenntou/modev-backend/domain/repository"
	"github.com/originbenntou/modev-backend/domain/vo"
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
		return nil
	}

	addDate, err := vo.NewDateTime(m.AddDate)
	if err != nil {
		return nil
	}

	url, err := vo.NewURL(m.Url)
	if err != nil {
		return nil
	}

	createdAt, err := vo.NewDateTime(m.CreatedAt)
	if err != nil {
		return nil
	}

	updatedAt, err := vo.NewDateTime(m.UpdatedAt)
	if err != nil {
		return nil
	}

	return &entity.TweetEntity{
		Id:        m.Id,
		Category:  *category,
		AddDate:   addDate.ToDateString(),
		Url:       *url,
		Tags:      nil,
		CreatedAt: *createdAt,
		UpdatedAt: *updatedAt,
	}
}
