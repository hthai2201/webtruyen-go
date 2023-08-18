package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

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
	}
}

type crawlStoryRequest struct {
	URL string `json:"url"       binding:"required"  example:"https://truyenfull.vn/tu-cam-270192"`
}

func (r *crawlerRoutes) crawlStory(c *gin.Context) {
	var request crawlStoryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - crawlStory")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	crawler, err := r.t.CrawlStory(
		c.Request.Context(), request.URL,
	)
	if err != nil {
		r.l.Error(err, "http - v1 - crawlStory")
		errorResponse(c, http.StatusInternalServerError, "crawler service problems")

		return
	}

	c.JSON(http.StatusOK, crawler)
}

type CrawlChaptersRequest struct {
	URL string `json:"url"       binding:"required"  example:"https://truyenfull.vn/tu-cam-270192"`
}

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

	c.JSON(http.StatusOK, gin.H{
		"total":           len(chapters),
		"newest_chapters": chapters[:10],
	})
}
