// internal/campaign/service_test.go
package campaign

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// mockRepository é a nossa implementacao falsa da interface Repository para os testes.
type mockRepository struct {
	mock.Mock
}

func (m *mockRepository) Create(campaign *Campaign) error {

	args := m.Called(campaign)
	return args.Error(0)
}

func (m *mockRepository) FindAll() ([]Campaign, error) {

	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]Campaign), args.Error(1)
}

func (m *mockRepository) FindByID(id int) (*Campaign, error) {

	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*Campaign), args.Error(1)
}

func (m *mockRepository) Update(campaign *Campaign) (*Campaign, error) {

	args := m.Called(campaign)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*Campaign), args.Error(1)
}

func (m *mockRepository) Delete(id int) error {

	args := m.Called(id)
	return args.Error(0)
}


var (
	someError        = errors.New("some error")
	expectedCampaign = &Campaign{
		ID:        1,
		Name:      "Campanha Teste",
		Budget:    1000,
		Status:    "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
)

// Teste para CreateCampaign
func TestService_CreateCampaign_Success(t *testing.T) {

	mockRepo := new(mockRepository)
	campaignService := NewService(mockRepo)
	campaignToCreate := &Campaign{Name: "Nova Campanha"}

	mockRepo.On("Create", campaignToCreate).Return(nil)

	err := campaignService.CreateCampaign(campaignToCreate)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

// Teste para GetAllCampaigns
func TestService_GetAllCampaigns_Success(t *testing.T) {
	// Arrange
	mockRepo := new(mockRepository)
	campaignService := NewService(mockRepo)
	expectedCampaigns := []Campaign{*expectedCampaign}

	mockRepo.On("FindAll").Return(expectedCampaigns, nil)

	// Act
	campaigns, err := campaignService.GetAllCampaigns()

	// Assert
	assert.NoError(t, err)
	assert.Len(t, campaigns, 1)
	assert.Equal(t, expectedCampaigns[0].Name, campaigns[0].Name)
	mockRepo.AssertExpectations(t)
}

// Teste para GetCampaignByID
func TestService_GetCampaignByID_Success(t *testing.T) {
	// Arrange
	mockRepo := new(mockRepository)
	campaignService := NewService(mockRepo)

	mockRepo.On("FindByID", expectedCampaign.ID).Return(expectedCampaign, nil)

	// Act
	campaign, err := campaignService.GetCampaignByID(expectedCampaign.ID)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, campaign)
	assert.Equal(t, expectedCampaign.Name, campaign.Name)
	mockRepo.AssertExpectations(t)
}

func TestService_GetCampaignByID_NotFound(t *testing.T) {
	// Arrange
	mockRepo := new(mockRepository)
	campaignService := NewService(mockRepo)
	
	mockRepo.On("FindByID", 99).Return(nil, sql.ErrNoRows)

	// Act
	campaign, err := campaignService.GetCampaignByID(99)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, campaign)
	assert.Equal(t, sql.ErrNoRows, err)
	mockRepo.AssertExpectations(t)
}

// Teste para UpdateCampaign
func TestService_UpdateCampaign_Success(t *testing.T) {
	// Arrange
	mockRepo := new(mockRepository)
	campaignService := NewService(mockRepo)
	campaignToUpdate := &Campaign{ID: 1, Name: "Campanha Atualizada"}

	mockRepo.On("Update", campaignToUpdate).Return(campaignToUpdate, nil)

	// Act
	updatedCampaign, err := campaignService.UpdateCampaign(1, campaignToUpdate)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, updatedCampaign)
	assert.Equal(t, "Campanha Atualizada", updatedCampaign.Name)
	mockRepo.AssertExpectations(t)
}

// Teste para DeleteCampaign
func TestService_DeleteCampaign_Success(t *testing.T) {
	// Arrange
	mockRepo := new(mockRepository)
	campaignService := NewService(mockRepo)
	
	mockRepo.On("Delete", 1).Return(nil)
	
	// Act
	err := campaignService.DeleteCampaign(1)
	
	// Assert
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestService_DeleteCampaign_NotFound(t *testing.T) {
	// Arrange
	mockRepo := new(mockRepository)
	campaignService := NewService(mockRepo)

	mockRepo.On("Delete", 99).Return(sql.ErrNoRows)

	// Act
	err := campaignService.DeleteCampaign(99)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, sql.ErrNoRows, err)
	mockRepo.AssertExpectations(t)
}
