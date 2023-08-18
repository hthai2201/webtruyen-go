package crawlerusecase

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/gocolly/colly"
	"github.com/hthai2201/webtruyen-go/internal/entity"
	"github.com/hthai2201/webtruyen-go/pkg/common"
)

// CrawlerUseCase -.
type CrawlerUseCase struct {
	repo   CrawlerRepo
	webAPI TruyenFullWebAPI
}

// New -.
func New(r CrawlerRepo, w TruyenFullWebAPI) *CrawlerUseCase {
	return &CrawlerUseCase{
		repo:   r,
		webAPI: w,
	}
}

func (uc *CrawlerUseCase) CrawlStory(ctx context.Context, url string) (entity.Story, error) {
	resultStory := entity.Story{}
	// slug
	resultStory.Slug = common.GetLastSlashValue(url)

	var wg sync.WaitGroup
	wg.Add(1)
	c := colly.NewCollector()
	c.OnHTML(".col-truyen-main", func(e *colly.HTMLElement) {
		resultStory.TID = e.ChildAttr(".col-truyen-main #truyen-id", "value")
		resultStory.Name = e.ChildText(".col-truyen-main h3.title")
		resultStory.Thumbnail = e.ChildAttr(".col-truyen-main .book img", "src")
		resultStory.Author.Name = e.ChildText(".col-truyen-main [itemprop='author']")
		resultStory.Author.Slug = common.GetLastSlashValue(e.ChildAttr(".col-truyen-main [itemprop='author']", "href"))

		e.ForEach(".col-truyen-main a[itemprop='genre']", func(i int, a *colly.HTMLElement) {
			c := entity.Category{
				Slug: common.GetLastSlashValue(a.Attr("href")),
				Name: a.Text,
			}
			resultStory.Categories = append(resultStory.Categories, c)
		})
		resultStory.Source = e.ChildText(".col-truyen-main .source")
		resultStory.Status = e.ChildText(".col-truyen-main .info >div:last-child span")
		description, descriptionError := e.DOM.Find(".col-truyen-main [itemprop='description']").First().Html()
		if descriptionError == nil {
			resultStory.Description = description
		}

		strRateValue, strRateValueError := strconv.ParseFloat(e.ChildText(".col-truyen-main [itemprop='ratingValue']"), 32)
		if strRateValueError == nil {
			resultStory.Rate.Value = float32(strRateValue)
		}

		strRateBestValue, strRateBestValueError := strconv.ParseFloat(e.ChildText(".col-truyen-main [itemprop='bestRating']"), 32)
		if strRateBestValueError == nil {
			resultStory.Rate.BestValue = float32(strRateBestValue)
		}
		strRateCount, strRateCountError := strconv.Atoi(e.ChildText(".col-truyen-main [itemprop='ratingCount']"))
		if strRateCountError == nil {
			resultStory.Rate.Count = strRateCount
		}

		wg.Done()
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Crawling Story... ", r.URL)
	})

	c.Visit(url)
	wg.Wait()
	uc.repo.StoreStory(ctx, &resultStory)
	return resultStory, nil
}
func (uc *CrawlerUseCase) CrawlChapters(ctx context.Context, url string) ([]entity.Chapter, error) {
	story := entity.Story{
		Slug: common.GetLastSlashValue(url),
	}
	err := uc.repo.FindStory(ctx, story, &story)
	if err != nil {
		story, err = uc.CrawlStory(ctx, url)
		if err != nil {
			return nil, fmt.Errorf("CrawlerUseCase - CrawlChapters - s.repo.FindStory: %w", err)

		}
	}
	rawHTMLChapters, err := uc.webAPI.GetStoryChapters(story, map[string]string{
		"page":   "1",
		"totalp": "1",
	})
	log.Println(rawHTMLChapters)
	chapters := []entity.Chapter{}
	return chapters, nil
}
