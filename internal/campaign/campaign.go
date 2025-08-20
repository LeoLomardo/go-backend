package campaign

import "time"

type Campaign struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Budget    float64   `json:"budget"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
