package crawlerusecase

import (
	"context"

	"github.com/hthai2201/webtruyen-go/internal/entity"
)

// CrawlerUseCase -.
type CrawlerUseCase struct {
	repo CrawlerRepo
}

// New -.
func New(r CrawlerRepo) *CrawlerUseCase {
	return &CrawlerUseCase{
		repo: r,
	}
}

func (uc *CrawlerUseCase) CrawlStory(context.Context) (entity.Story, error)
func (uc *CrawlerUseCase) CrawlStoryRate(context.Context) (entity.StoryRate, error)
func (uc *CrawlerUseCase) CrawlStoryList(context.Context) (entity.StoryList, error)
func (uc *CrawlerUseCase) CrawlCategory(context.Context) (entity.Category, error)
func (uc *CrawlerUseCase) CrawlChapter(context.Context) (entity.Chapter, error)
