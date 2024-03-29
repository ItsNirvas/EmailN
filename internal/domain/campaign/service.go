package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internalErrors"
)

type Service interface {
	Create(newCampaign contract.NewCampaign) (string, error)
	GetBy(ID string) (*contract.CampaignResponse, error)
}

type ServiceImp struct {
	RepositoryS Repository
}

func (s *ServiceImp) Create(newCampaign contract.NewCampaign) (string, error) {

	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Contacts)
	if err != nil {
		return "", err
	} else {
		err := s.RepositoryS.Save(campaign)
		if err != nil {
			return "", internalerrors.ErrInternal
		}
	}

	return campaign.ID, nil
}

func (s *ServiceImp) GetBy(id string) (*contract.CampaignResponse, error) {
	campaign, err := s.RepositoryS.GetBy(id)

	if err != nil {
		return nil, internalerrors.ErrInternal
	} else {
		return &contract.CampaignResponse{
			ID:      campaign.ID,
			Name:    campaign.Name,
			Content: campaign.Content,
			Status:  campaign.Status,
		}, nil
	}
}
