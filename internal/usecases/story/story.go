package storyusecase

import (
	"context"
	"fmt"

	"github.com/hthai2201/webtruyen-go/internal/entity"
	crawlerusecase "github.com/hthai2201/webtruyen-go/internal/usecases/crawler"
	"github.com/hthai2201/webtruyen-go/pkg/common"
)

// StoryUseCase -.
type StoryUseCase struct {
	repo    StoryRepo
	crawler crawlerusecase.Crawler
}

// New -.
func New(r StoryRepo, c crawlerusecase.Crawler) *StoryUseCase {
	return &StoryUseCase{
		repo:    r,
		crawler: c,
	}
}

func (uc *StoryUseCase) FindStory(ctx context.Context, slug string) (entity.Story, error) {
	story := entity.Story{
		Slug: slug,
	}
	err := uc.repo.FindStory(ctx, story, &story)
	if err != nil {
		story, err = uc.crawler.CrawlStoryBySlug(ctx, slug)
		if err != nil {
			return story, fmt.Errorf("StoryUseCase - FindStory - s.repo.FindStory: %w", err)
		}

	}

	fmt.Println(story)
	return story, nil
}
func (uc *StoryUseCase) FindStories(ctx context.Context, f entity.Story, p *common.Pagination) ([]entity.Story, error) {
	var resultStories []entity.Story
	err := uc.repo.FindStories(ctx, f, p, &resultStories)
	if err != nil {
		return resultStories, fmt.Errorf("StoryUseCase - FindStories - s.repo.FindStories: %w", err)

	}
	return resultStories, nil
}

func (uc *StoryUseCase) FindStoryChapter(ctx context.Context, sSlug string, cSlug string) (entity.Chapter, error) {
	var chapter entity.Chapter
	story := entity.Story{
		Slug: sSlug,
	}
	err := uc.repo.FindStory(ctx, story, &story)
	if err != nil {
		story, err = uc.crawler.CrawlStoryBySlug(ctx, sSlug)
		if err != nil {
			return chapter, fmt.Errorf("StoryUseCase - FindStoryChapter - s.repo.CrawlStoryBySlug: %w", err)
		}

	}
	err = uc.repo.FindStoryChapter(ctx, story.ID, cSlug, &chapter)
	if err != nil {
		chapter, err = uc.crawler.CrawlChapterBySlug(ctx, sSlug, cSlug)
		if err != nil {

			return chapter, fmt.Errorf("StoryUseCase - CrawlChapterBySlug: %w", err)
		}
	}
	return chapter, nil
}
