// Package usecase implements application business logic. Each logic group in own file.
package storyusecase

import (
	"context"

	"github.com/hthai2201/webtruyen-go/internal/entity"
	"github.com/hthai2201/webtruyen-go/pkg/common"
)

type (
	// Story -.
	Story interface {
		FindStory(ctx context.Context, url string) (entity.Story, error)
		FindStories(ctx context.Context, f entity.Story, p *common.Pagination) ([]entity.Story, error)
		FindStoryChapter(ctx context.Context, sSlug string, cSlug string) (entity.Chapter, error)
	}

	// StoryRepo -.
	StoryRepo interface {
		FindStory(context.Context, entity.Story, *entity.Story) error
		FindStories(ctx context.Context, f entity.Story, p *common.Pagination, t *[]entity.Story) error
		FindStoryChapter(ctx context.Context, storyID int, slug string, c *entity.Chapter) error
	}
)
