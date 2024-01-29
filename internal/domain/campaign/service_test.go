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

var (
	newCampaign = contract.NewCampaign{
		Name:     "Test X",
		Content:  "Body",
		Contacts: []string{"test1@gmail,com", "test2@gmail.com", "test3@gmail.com"},
	}

	service = Service{}
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
	newCampaign.Name = ""

	_, err := service.Create(newCampaign)

	assert.NotNil(err)
	assert.Equal("all fields must be filled", err.Error())
}

func TestCreateCampaign_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repositoryM := new(repositoryMock)
	repositoryM.On("Save", mock.Anything).Return(errors.New("error to save on database"))
	service.RepositoryS = repositoryM

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(internalerrors.ErrInternal, err))
}
