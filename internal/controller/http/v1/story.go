package v1

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hthai2201/webtruyen-go/internal/entity"
	storyusecase "github.com/hthai2201/webtruyen-go/internal/usecases/story"
	"github.com/hthai2201/webtruyen-go/pkg/logger"
)

type storyRoutes struct {
	t storyusecase.Story
	l logger.Interface
}

func newStoryRoutes(handler *gin.RouterGroup, t storyusecase.Story, l logger.Interface) {
	r := &storyRoutes{t, l}

	h := handler.Group("/stories")
	{
		h.GET("/:slug", r.getStoryBySlug)
	}
}

type getStoryBySlugRequest struct {
	Slug string `url:"slug" binding:"required"  example:"tu-cam-270192"`
}
type getStoryBySlugResponse struct {
	Story entity.Story `json:"story"`
}

// @Summary     Get Story Details
// @Description Get Story Details by Slug
// @ID          get-story-by-slug
// @Tags        story
// @Accept      json
// @Produce     json
// @Success     200 {object} getStoryBySlugResponse
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /stories/{slug} [post]
func (r *storyRoutes) getStoryBySlug(c *gin.Context) {
	var slug = c.Param("slug")
	if slug == "" {
		r.l.Error(errors.New("Invalid Slug"), "http - v1 - crawlStory")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	story, err := r.t.FindStory(
		c.Request.Context(), slug,
	)
	if err != nil {
		r.l.Error(err, "http - v1 - getStoryBySlug")
		errorResponse(c, http.StatusInternalServerError, "story service problems")

		return
	}
	fmt.Println(story)
	c.JSON(http.StatusOK, getStoryBySlugResponse{Story: story})
}
