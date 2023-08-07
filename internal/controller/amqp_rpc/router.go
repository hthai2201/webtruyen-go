package amqprpc

import (
	"github.com/hthai2201/webtruyen-go/internal/usecase"
	"github.com/hthai2201/webtruyen-go/pkg/rabbitmq/rmq_rpc/server"
)

// NewRouter -.
func NewRouter(t usecase.Translation) map[string]server.CallHandler {
	routes := make(map[string]server.CallHandler)
	{
		newTranslationRoutes(routes, t)
	}

	return routes
}
