package role

import (
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/infrastructure"
)

// UserRoutes struct
type Route struct {
	logger     framework.Logger
	handler    infrastructure.Router
	controller *Controller
}

func NewRoute(
	logger framework.Logger,
	handler infrastructure.Router,
	controller *Controller,
) *Route {
	return &Route{
		handler:    handler,
		logger:     logger,
		controller: controller,
	}

}

// Setup user routes
func RegisterRoute(r *Route) {
	r.logger.Info("Setting up routes")

	api := r.handler.Group("/api")

	api.POST("/role", r.controller.CreateRole)
	api.GET("/role", r.controller.GetRoles)

}
