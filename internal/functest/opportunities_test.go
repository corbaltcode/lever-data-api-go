package functest

import (
	"context"
	"os"
	"testing"

	lever "github.com/corbaltcode/lever-data-api-go"
)

// This tests the opportunities read APIs against the live Lever API.
func TestOpportunities(t *testing.T) {
	apiKey := os.Getenv("LEVER_API_KEY")
	if apiKey == "" {
		t.Skip("LEVER_API_KEY not set")
		return
	}

	client := lever.NewClient(lever.WithAPIKey(apiKey))
	req := lever.NewListOpportunitiesRequest()
	req.Limit = 5
	ctx := context.Background()
	resp, err := client.ListOpportunities(ctx, req)
	if err != nil {
		t.Fatal("Failed to list opportunities:", err)
	}

	if len(resp.Opportunities) == 0 {
		t.Fatal("No opportunities found")
	}

	for _, opportunity := range resp.Opportunities {
		t.Logf("%-36s %-48s", opportunity.ID, opportunity.Name)
		t.Logf("    Headline: %s", opportunity.Headline)
		t.Logf("    Location: %s", opportunity.Location)
		t.Logf("    ContactID: %s", opportunity.ContactID)
		t.Logf("    StageID: %s", opportunity.StageID)
	}

	t.Logf("HasNext: %v", resp.HasNext)
	t.Logf("Next: %s", resp.Next)
}

// This tests the opportunities read APIs against the live Lever API, with the stage expanded.
func TestOpportunitiesExpandStage(t *testing.T) {
	apiKey := os.Getenv("LEVER_API_KEY")
	if apiKey == "" {
		t.Skip("LEVER_API_KEY not set")
		return
	}

	client := lever.NewClient(lever.WithAPIKey(apiKey))
	req := lever.NewListOpportunitiesRequest()
	req.Limit = 5
	req.Expand = append(req.Expand, "stage")
	ctx := context.Background()
	resp, err := client.ListOpportunities(ctx, req)
	if err != nil {
		t.Fatal("Failed to list opportunities:", err)
	}

	if len(resp.Opportunities) == 0 {
		t.Fatal("No opportunities found")
	}

	for _, opportunity := range resp.Opportunities {
		t.Logf("%-36s %-48s", opportunity.ID, opportunity.Name)
		t.Logf("    Headline: %s", opportunity.Headline)
		t.Logf("    Location: %s", opportunity.Location)
		t.Logf("    ContactID: %s", opportunity.ContactID)
		t.Logf("    StageID: %s", opportunity.StageID)
		if opportunity.Stage != nil {
			t.Logf("    Stage: %v", opportunity.Stage)
		}
	}

	t.Logf("HasNext: %v", resp.HasNext)
	t.Logf("Next: %s", resp.Next)
}
