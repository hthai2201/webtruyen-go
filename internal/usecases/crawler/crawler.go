package crawlerusecase

import (
	"context"
	"errors"
	"fmt"
	URL "net/url"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
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

func (uc *CrawlerUseCase) CrawlStoryBySlug(ctx context.Context, slug string) (entity.Story, error) {
	return uc.CrawlStory(ctx, uc.webAPI.GetStoryURLBySlug(slug))

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
	go uc.CrawlChapters(ctx, url)
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

	chapters := []entity.Chapter{}
	nextChapterUrl, err := URL.JoinPath(url, "chuong-1")

	for nextChapterUrl != "" {
		var newChapter entity.Chapter
		newChapter, nextChapterUrl, err = parseChapter(nextChapterUrl, story.ID)
		fmt.Println(err, nextChapterUrl)
		if err != nil {
			fmt.Printf("Error visiting %s: %s\n", url, err)
			break
		}
		uc.repo.StoreChapter(ctx, &newChapter)
		chapters = append(chapters, newChapter)
		time.Sleep(200)
	}
	return chapters, nil
}

func parseChapter(url string, storyId int) (entity.Chapter, string, error) {
	newChapter := entity.Chapter{
		StoryID: storyId,
	}
	nextChapterUrl := ""
	var err error
	c := colly.NewCollector()
	isNotFound := false
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error when crawling... ", r.Request.URL.String())
	})
	c.OnHTML("#chapter-big-container", func(e *colly.HTMLElement) {
		if e == nil {
			isNotFound = true
		}
		newChapter.CID = e.ChildAttr("#chapter-id", "value")
		namePattern := `chương ([0-9-]+):?(.+)?`
		re := regexp.MustCompile("(?i)" + namePattern)
		matches := re.FindStringSubmatch(e.ChildText(".chapter-title"))
		if len(matches) >= 2 {
			newChapter.Name = strings.TrimSpace(matches[2])
			newChapter.Slug = common.GetLastSlashValue(e.ChildAttr(".chapter-title", "href"))

		}
		contentHtml, err := e.DOM.Find(".chapter-c").First().Html()
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(contentHtml))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		// Remove div elements and get the HTML of the remaining content
		doc.Find("div").Remove()
		newChapter.Content, err = doc.Find("body").Html()

		nextChapterUrl = e.ChildAttr("#next_chap", "href")
		if !strings.HasPrefix(nextChapterUrl, "http") {
			//last chapter
			nextChapterUrl = ""
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Crawling... ", r.URL.String())
	})
	err = c.Visit(url)

	if err != nil {
		return newChapter, nextChapterUrl, err
	}
	if isNotFound {
		return newChapter, nextChapterUrl, errors.New("Chapter not found")

	}
	return newChapter, nextChapterUrl, nil
}
