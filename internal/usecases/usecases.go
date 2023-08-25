package usecases

import (
	crawlerusecase "github.com/hthai2201/webtruyen-go/internal/usecases/crawler"
	crawlerpepo "github.com/hthai2201/webtruyen-go/internal/usecases/crawler/repo"
	truyenfullwebapi "github.com/hthai2201/webtruyen-go/internal/usecases/crawler/webapi"
	storyusecase "github.com/hthai2201/webtruyen-go/internal/usecases/story"
	storyrepo "github.com/hthai2201/webtruyen-go/internal/usecases/story/repo"
	translationusecase "github.com/hthai2201/webtruyen-go/internal/usecases/translation"
	translationrepo "github.com/hthai2201/webtruyen-go/internal/usecases/translation/repo"
	translationwebapi "github.com/hthai2201/webtruyen-go/internal/usecases/translation/webapi"
	"github.com/hthai2201/webtruyen-go/pkg/appctx"
)

type Usecases struct {
	Translation *translationusecase.TranslationUseCase
	Crawler     *crawlerusecase.CrawlerUseCase
	Story       *storyusecase.StoryUseCase
}

// New -.
func New(appCtx *appctx.AppContext) *Usecases {

	return &Usecases{
		Translation: translationusecase.New(translationrepo.New(appCtx), translationwebapi.New()),
		Crawler:     crawlerusecase.New(crawlerpepo.New(appCtx), truyenfullwebapi.New()),
		Story:       storyusecase.New(storyrepo.New(appCtx), crawlerusecase.New(crawlerpepo.New(appCtx), truyenfullwebapi.New())),
	}
}
