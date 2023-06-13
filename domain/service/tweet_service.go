package service

import (
	"context"
	"github.com/labstack/gommon/log"
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
	tweetWithTagModels, err := s.TweetRepository.FindByCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	var tweetEntities []*entity.TweetEntity
	for _, m := range tweetWithTagModels {
		if tweet, ok := findEntityById(tweetEntities, m.Id); ok {
			// nullable対応
			if m.Tag.Valid {
				tweet.Tags = append(tweet.Tags, m.Tag.String)
			}
		} else {
			tweetEntities = append(tweetEntities, modelToEntity(m))
		}
	}

	return tweetEntities, nil
}

func findEntityById(list []*entity.TweetEntity, id uint64) (*entity.TweetEntity, bool) {
	for _, e := range list {
		if e.Id == id {
			return e, true
		}
	}
	return nil, false
}

func modelToEntity(m *model.TweetWithTagModel) *entity.TweetEntity {
	category, err := vo.NewCategory(m.Category)
	if err != nil {
		log.Warn(err)
		return nil
	}

	url, err := vo.NewURL(m.Url)
	if err != nil {
		log.Warn(err)
		return nil
	}

	tags := make([]string, 0)
	if m.Tag.Valid {
		tags = append(tags, m.Tag.String)
	}

	createdAt, err := vo.NewDateTime(m.CreatedAt)
	if err != nil {
		log.Warn(err)
		return nil
	}

	updatedAt, err := vo.NewDateTime(m.UpdatedAt)
	if err != nil {
		log.Warn(err)
		return nil
	}

	return &entity.TweetEntity{
		Id:        m.Id,
		Category:  *category,
		AddDate:   m.AddDate,
		Url:       *url,
		Tags:      tags,
		CreatedAt: *createdAt,
		UpdatedAt: *updatedAt,
	}
}
