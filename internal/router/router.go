package router

import (
	"go-backend/internal/campaign"
	"go-backend/internal/middleware"
	"go-backend/internal/user"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Dependencias da campanha
	campaignRepo := campaign.NewRepository()
	campaignService := campaign.NewService(campaignRepo)
	campaignHandler := campaign.NewHandler(campaignService)

	
	app.Post("/login", user.Login)

	// Agrupa as rotas de campanha e aplica o middleware de autenticação
	campaignRoutes := app.Group("/campaigns", middleware.AuthMiddleware)

	// Rotas para Campanhas (CRUD)
	campaignRoutes.Post("/", campaignHandler.CreateCampaign)
	campaignRoutes.Get("/", campaignHandler.GetCampaigns)
	campaignRoutes.Get("/:id", campaignHandler.GetCampaign)
	campaignRoutes.Put("/:id", campaignHandler.UpdateCampaign)
	campaignRoutes.Delete("/:id", campaignHandler.DeleteCampaign)
}