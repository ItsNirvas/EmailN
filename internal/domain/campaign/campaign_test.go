package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign One"
	content  = "Shirt"
	contacts = []string{"email1@gmail.com", "email2@gmail.com"}
	fake     = faker.New()
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

func TestNewCampaign_ErrorOnMinName(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts)

	assert.Equal("name is required with minimum of 3 chars", err.Error())
}

func TestNewCampaign_ErrorOnMaxName(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(fake.Lorem().Text(30), content, contacts)

	assert.Equal("name is required with maximum of 24 chars", err.Error())
}

func TestNewCampaign_ErrorOnMinContent(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", contacts)

	assert.Equal("content is required with minimum of 3 chars", err.Error())
}

func TestNewCampaign_ErrorOnMaxContent(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, fake.Lorem().Text(1100), contacts)

	assert.Equal("content is required with maximum of 1024 chars", err.Error())
}

func TestNewCampaign_ErrorOnMinContacts(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{})

	assert.Equal("contacts is required with minimum of 1 email", err.Error())
}

func TestNewCampaign_ErrorOnInvalidContacts(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{"email_invalid"})

	assert.Equal("email is invalid", err.Error())
}
