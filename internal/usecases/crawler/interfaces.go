// Package usecase implements application business logic. Each logic group in own file.
package crawlerusecase

import (
	"context"

	"github.com/hthai2201/webtruyen-go/internal/entity"
	truyenfullwebapi "github.com/hthai2201/webtruyen-go/internal/usecases/crawler/webapi"
)

type (
	// Crawler -.
	Crawler interface {
		CrawlStory(ctx context.Context, url string) (entity.Story, error)
		CrawlStoryBySlug(ctx context.Context, slug string) (entity.Story, error)
		CrawlChapters(ctx context.Context, url string) ([]entity.Chapter, error)
		CrawlChapter(ctx context.Context, url string) (entity.Chapter, error)
		CrawlChapterBySlug(ctx context.Context, sSlug string, cSlug string) (entity.Chapter, error)
	}

	// CrawlerRepo -.
	CrawlerRepo interface {
		StoreStory(context.Context, *entity.Story) error
		StoreStoryRate(context.Context, *entity.StoryRate) error
		StoreStoryList(context.Context, *entity.StoryList) error
		StoreCategory(context.Context, *entity.Category) error
		StoreChapter(context.Context, *entity.Chapter) error
		StoreChapters(context.Context, *[]entity.Chapter) error
		StoreAuthor(context.Context, *entity.Author) error
		FindStory(context.Context, entity.Story, *entity.Story) error
	}
	TruyenFullWebAPI interface {
		GetStoryChapters(story entity.Story, opt map[string]string) (truyenfullwebapi.GetStoryChaptersResponse, error)
		GetStoryURLBySlug(slug string) string
		GetChapterURLBySlug(sSlug string, cSlug string) string
		ExtractStorySlug(url string) string
	}
)
