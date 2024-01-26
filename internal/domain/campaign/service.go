package campaign

import (
	"emailn/internal/contract"
)

type Service struct {
	RepositoryN Repository
}

func (s *Service) Create(newCampaign contract.NewCampaign) (string, error) {
	campaign, _ := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Contacts)
	s.RepositoryN.Save(campaign)

	return campaign.ID, nil
}
