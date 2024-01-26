package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign One"
	content  = "Shirt"
	contacts = []string{"email1@gmail.com", "email2@gmail.com"}
)

func TestNewCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
	assert.Equal(campaign.Name, name)
	assert.Greater(campaign.Created, time.Now().Add(-time.Minute))
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}

func TestNewCampaign_ErrorOnNilFields(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", "", []string{})

	assert.Equal("all fields must be filled", err.Error())
}
