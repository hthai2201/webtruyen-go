package storyrepo

import (
	"context"
	"fmt"

	"github.com/hthai2201/webtruyen-go/internal/entity"
	"github.com/hthai2201/webtruyen-go/pkg/appctx"
)

const _defaultEntityCap = 64

// StoryRepo -.
type StoryRepo struct {
	appCtx *appctx.AppContext
}

// New -.
func New(appCtx *appctx.AppContext) *StoryRepo {
	return &StoryRepo{appCtx}
}

func (r *StoryRepo) FindStory(ctx context.Context, f entity.Story, t *entity.Story) error {
	db := r.appCtx.GetDBConnection()
	result := db.Table(entity.Story{}.TableName()).Where(f).First(t)
	if result.Error != nil {
		return fmt.Errorf("StoryRepo - FindStory - r.Pool.Exec: %w", result.Error)
	}

	return nil
}
