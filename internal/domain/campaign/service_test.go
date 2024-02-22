package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internalErrors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *repositoryMock) Get() ([]Campaign, error) {
	//args := r.Called(campaign)
	return nil, nil
}

func (r *repositoryMock) GetByID(ID string) (*Campaign, error) {
	//args := r.Called(campaign)
	return nil, nil
}

var (
	newCampaign = contract.NewCampaign{
		Name:     "Test X",
		Content:  "Shirt",
		Contacts: []string{"test1@gmail.com", "test2@gmail.com", "test3@gmail.com"},
	}

	service = ServiceImp{}
)

func TestCreateCampaign(t *testing.T) {

	repositoryM := new(repositoryMock)
	repositoryM.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name || campaign.Content != newCampaign.Content || len(campaign.Contacts) != len(newCampaign.Contacts) {
			return false
		} else {
			return true
		}
	})).Return(nil)
	service.RepositoryS = repositoryM

	service.Create(newCampaign)

	repositoryM.AssertExpectations(t)
}

func TestCreateCampaign_ValidadeErrors(t *testing.T) {
	assert := assert.New(t)

	_, err := service.Create(contract.NewCampaign{})

	assert.False(errors.Is(internalerrors.ErrInternal, err))
}

func TestCreateCampaign_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repositoryM := new(repositoryMock)
	repositoryM.On("Save", mock.Anything).Return(errors.New("error to save on database"))
	service.RepositoryS = repositoryM

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(internalerrors.ErrInternal, err))
}
