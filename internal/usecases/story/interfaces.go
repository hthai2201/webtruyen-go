// Package usecase implements application business logic. Each logic group in own file.
package storyusecase

import (
	"context"

	"github.com/hthai2201/webtruyen-go/internal/entity"
)

type (
	// Story -.
	Story interface {
		FindStory(ctx context.Context, url string) (entity.Story, error)
	}

	// StoryRepo -.
	StoryRepo interface {
		FindStory(context.Context, entity.Story, *entity.Story) error
	}
)
