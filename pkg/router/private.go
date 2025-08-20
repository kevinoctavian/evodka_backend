package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kevinoctavian/evodka_backend/app/controller"
	"github.com/kevinoctavian/evodka_backend/pkg/middleware"
)

func PrivateRoutes(app *fiber.App) {
	route := app.Group("/api/v1", middleware.JWTProtected())

	// Elections router
	electionsRoute := route.Group("/elections")
	electionsRoute.Get("/", controller.GetElections)
	electionsRoute.Get("/:id", controller.GetElection)

	// router for manage elections by admin only
	electionsRouteAdminOnly := electionsRoute.Group("", middleware.IsAdmin)
	electionsRouteAdminOnly.Post("/", controller.CreateElection)
	electionsRouteAdminOnly.Put("/:id", controller.UpdateElection)
	electionsRouteAdminOnly.Delete("/:id", controller.DeleteElection)

	// voters router
	voterRoute := route.Group("/voters")
	voterRoute.Get("/", controller.GetVoters)
	voterRoute.Post("/", controller.CreateVoter)
	voterRoute.Delete("/:id", controller.DeleteVoter)

	// users router
	route.Get("/users", controller.GetUsers)
	route.Put("/users/:id", controller.UpdateUser)
	route.Delete("/users/:id", controller.DeleteUser)
}
