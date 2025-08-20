// internal/campaign/handler.go
package campaign

import (
	"go-backend/internal/database"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateCampaign(c *fiber.Ctx) error {
	campaign := new(Campaign)
	if err := c.BodyParser(campaign); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	query := "INSERT INTO campaigns (name, budget, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	err := database.DB.QueryRow(query, campaign.Name, campaign.Budget, campaign.Status, time.Now(), time.Now()).Scan(&campaign.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(campaign)
}

// Listar Campanhas
func GetCampaigns(c *fiber.Ctx) error {
	query := "SELECT * FROM campaigns ORDER_BY created_at DESC"

	rows, err := database.DB.Query(query)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao buscar campanhas no banco de dados.",
		})
	}

	defer rows.Close()

	campaigns := []Campaign{}

	for rows.Next() {
		var campaign Campaign

		err := rows.Scan(&campaign.ID, &campaign.Name, &campaign.Budget, &campaign.Status, &campaign.CreatedAt, &campaign.UpdatedAt)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Erro ao processar os dados da campanha.",
			})
		}
		campaigns = append(campaigns, campaign)
	}

	return c.Status(fiber.StatusOK).JSON(campaigns)
}

// Buscar Campanha por ID [cite: 112]
func GetCampaign(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	// Implemente a logica para buscar uma campanha por ID
	return c.SendString("Get Campaign with ID: " + strconv.Itoa(id))
}

// Atualizar Campanha [cite: 114]
func UpdateCampaign(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	// Implemente a logica para atualizar a campanha
	// ...
	return c.SendString("Update Campaign with ID: " + strconv.Itoa(id))
}

// Deletar Campanha [cite: 116]
func DeleteCampaign(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	// Implemente a logica para deletar a campanha
	// ...
	return c.SendString("Delete Campaign with ID: " + strconv.Itoa(id))
}
