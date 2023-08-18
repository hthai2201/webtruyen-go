// Package usecase implements application business logic. Each logic group in own file.
package crawlerusecase

import (
	"context"

	"github.com/hthai2201/webtruyen-go/internal/entity"
)

type (
	// Crawler -.
	Crawler interface {
		CrawlStory(ctx context.Context, url string) (entity.Story, error)
	}

	// CrawlerRepo -.
	CrawlerRepo interface {
		StoreStory(context.Context, *entity.Story) error
		StoreStoryRate(context.Context, *entity.StoryRate) error
		StoreStoryList(context.Context, *entity.StoryList) error
		StoreCategory(context.Context, *entity.Category) error
		StoreChapter(context.Context, *entity.Chapter) error
		StoreAuthor(context.Context, *entity.Author) error
	}
)
