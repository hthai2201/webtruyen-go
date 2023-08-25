package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hthai2201/webtruyen-go/internal/entity"
	crawlerusecase "github.com/hthai2201/webtruyen-go/internal/usecases/crawler"
	"github.com/hthai2201/webtruyen-go/pkg/logger"
)

type crawlerRoutes struct {
	t crawlerusecase.Crawler
	l logger.Interface
}

func newCrawlerRoutes(handler *gin.RouterGroup, t crawlerusecase.Crawler, l logger.Interface) {
	r := &crawlerRoutes{t, l}

	h := handler.Group("/crawler")
	{
		h.POST("/crawl-story", r.crawlStory)
		h.POST("/crawl-chapters", r.crawlChapters)
		h.POST("/crawl-chapter", r.crawlChapter)
	}
}

type crawlStoryRequest struct {
	URL string `json:"url"       binding:"required"  example:"https://truyenfull.vn/tu-cam-270192"`
}
type crawlStoryResponse struct {
	Story entity.Story `json:"story"`
}

// @Summary     Crawl Story
// @Description Crawl a story from a URL
// @ID          crawl-story
// @Tags        crawler
// @Accept      json
// @Produce     json
// @Param       request body crawlStoryRequest true "Crawl story request"
// @Success     200 {object} crawlStoryResponse
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /crawler/crawl-story [post]
func (r *crawlerRoutes) crawlStory(c *gin.Context) {
	var request crawlStoryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - crawlStory")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	story, err := r.t.CrawlStory(
		c.Request.Context(), request.URL,
	)
	if err != nil {
		r.l.Error(err, "http - v1 - crawlStory")
		errorResponse(c, http.StatusInternalServerError, "crawler service problems")

		return
	}

	c.JSON(http.StatusOK, crawlStoryResponse{Story: story})
}

type CrawlChaptersRequest struct {
	URL string `json:"url"       binding:"required"  example:"https://truyenfull.vn/tu-cam-270192"`
}
type CrawlChaptersResponse struct {
	total           int
	newest_chapters []entity.Chapter
}

// @Summary     Crawl Chapters
// @Description Crawl chapters from a URL
// @ID          crawl-chapters
// @Tags        crawler
// @Accept      json
// @Produce     json
// @Param       request body CrawlChaptersRequest true "Crawl chapters request"
// @Success     200 {object} CrawlChaptersResponse
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /crawler/crawl-chapters [post]
func (r *crawlerRoutes) crawlChapters(c *gin.Context) {
	var request CrawlChaptersRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - CrawlChapters")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	chapters, err := r.t.CrawlChapters(
		c.Request.Context(), request.URL,
	)
	if err != nil {
		r.l.Error(err, "http - v1 - CrawlChapters")
		errorResponse(c, http.StatusInternalServerError, "crawler service problems")

		return
	}

	c.JSON(http.StatusOK, CrawlChaptersResponse{
		total:           len(chapters),
		newest_chapters: chapters,
	})
}

type CrawlChapterRequest struct {
	URL string `json:"url"       binding:"required"  example:"https://truyenfull.vn/tu-cam-270192"`
}
type CrawlChapterResponse struct {
	Chapter entity.Chapter `json:"chapter"`
}

// @Summary     Crawl Chapter
// @Description Crawl chapter from a URL
// @ID          crawl-chapter
// @Tags        crawler
// @Accept      json
// @Produce     json
// @Param       request body CrawlChapterRequest true "Crawl chapter request"
// @Success     200 {object} CrawlChapterResponse
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /crawler/crawl-chapters [post]
func (r *crawlerRoutes) crawlChapter(c *gin.Context) {
	var request CrawlChapterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - CrawlChapter")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	chapter, err := r.t.CrawlChapter(
		c.Request.Context(), request.URL,
	)
	if err != nil {
		r.l.Error(err, "http - v1 - CrawlChapters")
		errorResponse(c, http.StatusInternalServerError, "crawler service problems")

		return
	}

	c.JSON(http.StatusOK, CrawlChapterResponse{
		Chapter: chapter,
	})
}
