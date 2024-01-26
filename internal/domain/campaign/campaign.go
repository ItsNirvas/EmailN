package campaign

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string
}

type Campaign struct {
	ID       string
	Name     string
	Created  time.Time
	Content  string
	Contacts []Contact
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	if name == "" || content == "" || len(emails) == 0 {
		return nil, errors.New("all fields must be filled")
	}

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
