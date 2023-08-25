// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger docs.
	_ "github.com/hthai2201/webtruyen-go/docs"
	"github.com/hthai2201/webtruyen-go/internal/usecases"
	"github.com/hthai2201/webtruyen-go/pkg/logger"
)

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8000
// @BasePath    /v1
func NewRouter(handler *gin.Engine, l logger.Interface, t *usecases.Usecases) {

	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers
	h := handler.Group("/v1")
	{
		// newTranslationRoutes(h, t.Translation, l)
		newCrawlerRoutes(h, t.Crawler, l)
		newStoryRoutes(h, t.Story, l)
	}
}
