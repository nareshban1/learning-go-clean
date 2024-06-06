package user

import (
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/infrastructure"
	"clean-architecture/pkg/middlewares"
)

// UserRoutes struct
type Route struct {
	logger              framework.Logger
	handler             infrastructure.Router
	controller          *Controller
	authMiddleware      middlewares.CognitoMiddleWare
	uploadMiddleware    middlewares.UploadMiddleware
	rateLimitMiddleware middlewares.RateLimitMiddleware
}

func NewRoute(
	logger framework.Logger,
	handler infrastructure.Router,
	controller *Controller,
	authMiddleware middlewares.CognitoAuthMiddleware,
	uploadMiddleware middlewares.UploadMiddleware,
	rateLimit middlewares.RateLimitMiddleware,
) *Route {
	return &Route{
		handler:             handler,
		logger:              logger,
		controller:          controller,
		authMiddleware:      authMiddleware,
		uploadMiddleware:    uploadMiddleware,
		rateLimitMiddleware: rateLimit,
	}

}

// Setup user routes
func RegisterRoute(r *Route) {
	r.logger.Info("Setting up routes")

	api := r.handler.Group("/api").Use(r.authMiddleware.Handle())

	api.POST("/user", r.controller.CreateUser)
	api.GET("/user/:id", r.controller.GetUserByID)

}
