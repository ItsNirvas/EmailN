package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internalErrors"
)

type Service struct {
	RepositoryS Repository
}

func (s *Service) Create(newCampaign contract.NewCampaign) (string, error) {

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
