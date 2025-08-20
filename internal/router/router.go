// router/router.go
package router

import (
	"go-backend/internal/campaign" // falta criar
	"go-backend/internal/middleware"
	"go-backend/internal/user"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/login", user.Login)

	campaignRoutes := app.Group("/campaigns", middleware.AuthMiddleware)

	// Rotas para Campanhas (CRUD)
	campaignRoutes.Post("/", campaign.CreateCampaign)      // POST /campaigns
	campaignRoutes.Get("/", campaign.GetCampaigns)         // GET /campaigns
	campaignRoutes.Get("/:id", campaign.GetCampaign)       // GET /campaigns/:id
	campaignRoutes.Put("/:id", campaign.UpdateCampaign)    // PUT /campaigns/:id
	campaignRoutes.Delete("/:id", campaign.DeleteCampaign) // DELETE /campaigns/:id
}
