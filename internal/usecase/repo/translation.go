package repo

import (
	"context"
	"fmt"

	"github.com/hthai2201/webtruyen-go/internal/entity"
	"github.com/hthai2201/webtruyen-go/pkg/appctx"
)

const _defaultEntityCap = 64

// TranslationRepo -.
type TranslationRepo struct {
	appCtx *appctx.AppContext
}

// New -.
func New(appCtx *appctx.AppContext) *TranslationRepo {
	return &TranslationRepo{appCtx}
}

// GetHistory -.
func (r *TranslationRepo) GetHistory(ctx context.Context) ([]entity.Translation, error) {
	db := r.appCtx.GetDBConnection()
	var history []entity.Translation
	result := db.Table(entity.Translation{}.GetTableName()).Find(&history)

	if result.Error != nil {
		return nil, fmt.Errorf("TranslationRepo - GetHistory - r.Pool.Query: %w", result.Error)
	}

	return history, nil
}

// Store -.
func (r *TranslationRepo) Store(ctx context.Context, t entity.Translation) error {
	db := r.appCtx.GetDBConnection()
	result := db.Table(entity.Translation{}.GetTableName()).Create(t)

	if result.Error != nil {
		return fmt.Errorf("TranslationRepo - Store - r.Pool.Exec: %w", result.Error)
	}

	return nil
}
