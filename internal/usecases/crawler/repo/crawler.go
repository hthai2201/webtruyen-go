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
func (r *CrawlerRepo) StoreStory(ctx context.Context, t entity.Story) error {
	db := r.appCtx.GetDBConnection()
	result := db.Table(entity.Story{}.GetTableName()).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(t)

	if result.Error != nil {
		return fmt.Errorf("CrawlerRepo - StoreStory - r.Pool.Exec: %w", result.Error)
	}

	return nil
}

func (r *CrawlerRepo) StoreCategory(ctx context.Context, t entity.Category) error {
	db := r.appCtx.GetDBConnection()
	result := db.Table(entity.Category{}.GetTableName()).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(t)

	if result.Error != nil {
		return fmt.Errorf("CrawlerRepo - StoreCategory - r.Pool.Exec: %w", result.Error)
	}

	return nil
}
func (r *CrawlerRepo) StoreStoryList(ctx context.Context, t entity.StoryList) error {
	db := r.appCtx.GetDBConnection()
	result := db.Table(entity.StoryList{}.GetTableName()).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(t)

	if result.Error != nil {
		return fmt.Errorf("CrawlerRepo - StoreStoryList - r.Pool.Exec: %w", result.Error)
	}

	return nil
}
func (r *CrawlerRepo) StoreStoryRate(ctx context.Context, t entity.StoryRate) error {
	db := r.appCtx.GetDBConnection()
	result := db.Table(entity.StoryRate{}.GetTableName()).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(t)

	if result.Error != nil {
		return fmt.Errorf("CrawlerRepo - StoreStoryList - r.Pool.Exec: %w", result.Error)
	}

	return nil
}
func (r *CrawlerRepo) StoreChapter(ctx context.Context, t entity.Chapter) error {
	db := r.appCtx.GetDBConnection()
	result := db.Table(entity.Chapter{}.GetTableName()).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(t)

	if result.Error != nil {
		return fmt.Errorf("CrawlerRepo - StoreChapter - r.Pool.Exec: %w", result.Error)
	}

	return nil
}
