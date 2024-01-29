package campaign

import (
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string `validade:"email"`
}

type Campaign struct {
	ID       string    `validade:"required"`
	Name     string    `validate:"min=3,max=24"`
	Created  time.Time `validade:"required"`
	Content  string    `validate:"min=3,max=1024"`
	Contacts []Contact `validate:"min=1,dive"`
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		contacts[index].Email = email
	}

	return &Campaign{
		ID:       xid.New().String(),
		Name:     name,
		Content:  content,
		Created:  time.Now(),
		Contacts: contacts,
	}, nil
}
