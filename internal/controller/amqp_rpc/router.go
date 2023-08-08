package amqprpc

import (
	"github.com/hthai2201/webtruyen-go/internal/usecases"
	"github.com/hthai2201/webtruyen-go/pkg/rabbitmq/rmq_rpc/server"
)

// NewRouter -.
func NewRouter(t *usecases.Usecases) map[string]server.CallHandler {
	routes := make(map[string]server.CallHandler)
	{
		newTranslationRoutes(routes, t.Translation)
	}

	return routes
}
