package campaign

type Service interface {
	CreateCampaign(campaign *Campaign) error
	GetAllCampaigns() ([]Campaign, error)
	GetCampaignByID(id int) (*Campaign, error)
	UpdateCampaign(id int, campaignData *Campaign) (*Campaign, error)
	DeleteCampaign(id int) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateCampaign(campaign *Campaign) error {
	// Aqui poderíamos adicionar lógica de negócio, como validações.
	return s.repo.Create(campaign)
}

func (s *service) GetAllCampaigns() ([]Campaign, error) {
	return s.repo.FindAll()
}

func (s *service) GetCampaignByID(id int) (*Campaign, error) {
	return s.repo.FindByID(id)
}

func (s *service) UpdateCampaign(id int, campaignData *Campaign) (*Campaign, error) {
	campaignData.ID = id
	return s.repo.Update(campaignData)
}

func (s *service) DeleteCampaign(id int) error {
	return s.repo.Delete(id)
}