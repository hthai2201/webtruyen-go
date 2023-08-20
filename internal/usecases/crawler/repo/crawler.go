package crawlerpepo

import (
	"context"
	"fmt"

	"github.com/hthai2201/webtruyen-go/internal/entity"
	"github.com/hthai2201/webtruyen-go/pkg/appctx"
	"gorm.io/gorm/clause"
)

const _defaultEntityCap = 64

// CrawlerRepo -.
type CrawlerRepo struct {
	appCtx *appctx.AppContext
}

// New -.
func New(appCtx *appctx.AppContext) *CrawlerRepo {
	return &CrawlerRepo{appCtx}
}

// Store -.
func (r *CrawlerRepo) StoreStory(ctx context.Context, t *entity.Story) error {
	db := r.appCtx.GetDBConnection()
	result := db.Table(entity.Story{}.TableName()).Create(t)

	if result.Error != nil {
		return fmt.Errorf("CrawlerRepo - StoreStory - r.Pool.Exec: %w", result.Error)
	}

	return nil
}

func (r *CrawlerRepo) StoreCategory(ctx context.Context, t *entity.Category) error {
	db := r.appCtx.GetDBConnection()
	result := db.Table(entity.Category{}.TableName()).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(t)

	if result.Error != nil {
		return fmt.Errorf("CrawlerRepo - StoreCategory - r.Pool.Exec: %w", result.Error)
	}

	return nil
}
func (r *CrawlerRepo) StoreStoryList(ctx context.Context, t *entity.StoryList) error {
	db := r.appCtx.GetDBConnection()
	result := db.Table(entity.StoryList{}.TableName()).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(t)

	if result.Error != nil {
		return fmt.Errorf("CrawlerRepo - StoreStoryList - r.Pool.Exec: %w", result.Error)
	}

	return nil
}
func (r *CrawlerRepo) StoreStoryRate(ctx context.Context, t *entity.StoryRate) error {
	db := r.appCtx.GetDBConnection()
	result := db.Table(entity.StoryRate{}.TableName()).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(t)

	if result.Error != nil {
		return fmt.Errorf("CrawlerRepo - StoreStoryList - r.Pool.Exec: %w", result.Error)
	}

	return nil
}
func (r *CrawlerRepo) StoreChapter(ctx context.Context, t *entity.Chapter) error {
	db := r.appCtx.GetDBConnection()
	result := db.Table(entity.Chapter{}.TableName()).Create(t)

	if result.Error != nil {
		return fmt.Errorf("CrawlerRepo - StoreChapter - r.Pool.Exec: %w", result.Error)
	}

	return nil
}
func (r *CrawlerRepo) StoreChapters(ctx context.Context, t *[]entity.Chapter) error {
	db := r.appCtx.GetDBConnection()
	result := db.Table(entity.Chapter{}.TableName()).Create(t)

	if result.Error != nil {
		return fmt.Errorf("CrawlerRepo - StoreChapters - r.Pool.Exec: %w", result.Error)
	}

	return nil
}
func (r *CrawlerRepo) StoreAuthor(ctx context.Context, t *entity.Author) error {
	db := r.appCtx.GetDBConnection()
	result := db.Table(entity.Author{}.TableName()).Create(t)
	if result.Error != nil {
		return fmt.Errorf("CrawlerRepo - StoreAuthor - r.Pool.Exec: %w", result.Error)
	}

	return nil
}
func (r *CrawlerRepo) FindStory(ctx context.Context, f entity.Story, t *entity.Story) error {
	db := r.appCtx.GetDBConnection()
	result := db.Table(entity.Story{}.TableName()).Where(f).First(t)
	if result.Error != nil {
		return fmt.Errorf("CrawlerRepo - FindStory - r.Pool.Exec: %w", result.Error)
	}

	return nil
}
