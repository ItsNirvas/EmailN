package campaign

import (
	"emailn/internal/contract"
	"testing"

	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func TestCreateCampaign(t *testing.T) {
	newCampaign := contract.NewCampaign{
		Name:     "Test X",
		Content:  "Body",
		Contacts: []string{"test1@gmail,com", "test2@gmail.com", "test3@gmail.com"},
	}
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name || campaign.Content != newCampaign.Content || len(campaign.Contacts) != len(newCampaign.Contacts) {
			return false
		} else {
			return true
		}
	})).Return(nil)
	service := Service{RepositoryN: repositoryMock}

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}
