package storyrepo

import (
	"context"
	"fmt"

	"github.com/hthai2201/webtruyen-go/internal/entity"
	"github.com/hthai2201/webtruyen-go/pkg/appctx"
	"github.com/hthai2201/webtruyen-go/pkg/common"
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
func (r *StoryRepo) FindStories(ctx context.Context, f entity.Story, p *common.Pagination, t *[]entity.Story) error {
	db := r.appCtx.GetDBConnection()
	p.Fulfill()
	result := db.Table(entity.Story{}.TableName()).
		Where(f).
		Offset((p.Page - 1) * p.Limit).
		Limit(p.Limit).
		Find(t)
	if result.Error != nil {
		return fmt.Errorf("StoryRepo - FindStories - r.Pool.Exec: %w", result.Error)
	}
	db.Table(entity.Story{}.TableName()).
		Where(f).Count(&p.Total)

	return nil
}
func (r *StoryRepo) FindStoryChapter(ctx context.Context, storyID int, slug string, c *entity.Chapter) error {
	db := r.appCtx.GetDBConnection()
	result := db.Table(entity.Chapter{}.TableName()).
		Where("story_id = ?", storyID).
		Where("slug =?", slug).First(c)
	if result.Error != nil {
		return fmt.Errorf("StoryRepo - FindStories - r.Pool.Exec: %w", result.Error)
	}

	return nil
}
