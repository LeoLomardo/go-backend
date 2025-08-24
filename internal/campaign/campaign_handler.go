package campaign

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateCampaign(c *fiber.Ctx) error {

	var campaign Campaign

	if err := c.BodyParser(&campaign); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ERROR": "corpo da requisição invalido"})
	}

	if err := h.service.CreateCampaign(&campaign); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"ERROR": "erro ao criar campanha"})
	}

	return c.Status(fiber.StatusCreated).JSON(campaign)
}

func (h *Handler) GetCampaigns(c *fiber.Ctx) error {

	campaigns, err := h.service.GetAllCampaigns()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"ERROR": "erro ao buscar campanhas"})
	}

	return c.Status(fiber.StatusOK).JSON(campaigns)
}

func (h *Handler) GetCampaign(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ERROR": "ID de campanha invalido"})
	}

	campaign, err := h.service.GetCampaignByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"ERROR": "campanha nao encontrada"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"ERROR": "erro ao buscar campanha"})
	}

	return c.Status(fiber.StatusOK).JSON(campaign)
}

func (h *Handler) UpdateCampaign(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ERROR": "ID de campanha invalido"})
	}

	var campaignData Campaign
	if err := c.BodyParser(&campaignData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ERROR": "corpo da requisição invalido"})
	}

	updatedCampaign, err := h.service.UpdateCampaign(id, &campaignData)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"ERROR": "campanha nao encontrada"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"ERROR": "erro ao atualizar campanha"})
	}

	return c.Status(fiber.StatusOK).JSON(updatedCampaign)
}

func (h *Handler) DeleteCampaign(c *fiber.Ctx) error {
	
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ERROR": "ID de campanha invalido"})
	}

	if err := h.service.DeleteCampaign(id); err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"ERROR": "campanha nao encontrada"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"ERROR": "erro ao deletar campanha"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}