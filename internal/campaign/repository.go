package campaign

import (
	"database/sql"
	"go-backend/internal/database"
	"time"
)

type Repository interface {
	Create(campaign *Campaign) error
	FindAll() ([]Campaign, error)
	FindByID(id int) (*Campaign, error)
	Update(campaign *Campaign) (*Campaign, error)
	Delete(id int) error
}

type postgresRepository struct{}

//cria nova instancia do repositorio.
func NewRepository() Repository {

	return &postgresRepository{}
}

func (r *postgresRepository) Create(campaign *Campaign) error {

	query := "INSERT INTO campaigns (name, budget, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at"

	return database.DB.QueryRow(query, campaign.Name, campaign.Budget, campaign.Status, time.Now(), time.Now()).Scan(&campaign.ID, &campaign.CreatedAt, &campaign.UpdatedAt)
}

func (r *postgresRepository) FindAll() ([]Campaign, error) {

	query := "SELECT id, name, budget, status, created_at, updated_at FROM campaigns ORDER BY id ASC"

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var campaigns []Campaign
	for rows.Next() {

		var c Campaign
		if err := rows.Scan(&c.ID, &c.Name, &c.Budget, &c.Status, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, err
		}
		
		campaigns = append(campaigns, c)
	}

	return campaigns, nil
}

func (r *postgresRepository) FindByID(id int) (*Campaign, error) {

	query := "SELECT id, name, budget, status, created_at, updated_at FROM campaigns WHERE id = $1"

	var c Campaign
	err := database.DB.QueryRow(query, id).Scan(&c.ID, &c.Name, &c.Budget, &c.Status, &c.CreatedAt, &c.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *postgresRepository) Update(campaign *Campaign) (*Campaign, error) {
	
	query := `UPDATE campaigns 
			  SET name = $1, budget = $2, status = $3, updated_at = NOW() 
			  WHERE id = $4
			  RETURNING updated_at`

	err := database.DB.QueryRow(
		query,
		campaign.Name,
		campaign.Budget,
		campaign.Status,
		campaign.ID,
	).Scan(&campaign.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return campaign, nil
}

func (r *postgresRepository) Delete(id int) error {

	query := "DELETE FROM campaigns WHERE id = $1"

	result, err := database.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows 
	}
	
	return nil
}