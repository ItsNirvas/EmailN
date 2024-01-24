package campaign

import "testing"

func TestNewCampaign(t *testing.T) {
	name := "Campaign One"
	content := "Body"
	contacts := []string{"email1@gmail.com", "email2@gmail.com"}

	campaign := NewCampaign(name, content, contacts)

	if campaign.ID != "1" {
		t.Errorf("ID Expected 1. Received: ", campaign.ID)
	} else if campaign.Name != name {
		t.Errorf("Name Expected: ", name, ". Received: ", campaign.Name)
	} else if campaign.Content != content {
		t.Errorf("Content Expected: ", content, ". Received: ", campaign.Content)
	} else if len(campaign.Contacts) != len(contacts) {
		t.Errorf("Contacts Expected: ", contacts, ". Received: ", campaign.Contacts)
	}
}
