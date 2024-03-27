package functest

import (
	"context"
	"os"
	"testing"

	lever "github.com/corbaltcode/lever-data-api-go"
)

// This tests the archive reasons against the live Lever API.
func TestArchiveReasons(t *testing.T) {
	apiKey := os.Getenv("LEVER_API_KEY")
	if apiKey == "" {
		t.Skip("LEVER_API_KEY not set")
		return
	}

	client := lever.NewClient(lever.WithAPIKey(apiKey))
	req := lever.NewListArchiveReasonsRequest()
	ctx := context.Background()
	resp, err := client.ListArchiveReasons(ctx, req)
	if err != nil {
		t.Fatal("Failed to list archive reasons:", err)
	}

	if len(resp.ArchiveReasons) == 0 {
		t.Fatal("No archive reasons found")
	}

	for _, reason := range resp.ArchiveReasons {
		t.Logf("%-36s %-24s %-8s %-9s", reason.ID, reason.Text, reason.Status, reason.Type)
	}

	for _, reason := range resp.ArchiveReasons {
		req := lever.NewGetArchiveReasonRequest(reason.ID)
		resp, err := client.GetArchiveReason(ctx, req)
		if err != nil {
			t.Errorf("Failed to get archive reason for reason id %s: %s", reason.ID, err)
			continue
		}

		if resp.ArchiveReason == nil {
			t.Errorf("No archive reason found for reason id %s", reason.ID)
			continue
		}

		if resp.ArchiveReason.ID != reason.ID {
			t.Errorf("Archive reason id mismatch: expected %s, got %s", reason.ID, resp.ArchiveReason.ID)
		}

		if resp.ArchiveReason.Text != reason.Text {
			t.Errorf("Archive reason text mismatch: expected %s, got %s", reason.Text, resp.ArchiveReason.Text)
		}

		if resp.ArchiveReason.Status != reason.Status {
			t.Errorf("Archive reason status mismatch: expected %s, got %s", reason.Status, resp.ArchiveReason.Status)
		}

		if resp.ArchiveReason.Type != reason.Type {
			t.Errorf("Archive reason type mismatch: expected %s, got %s", reason.Type, resp.ArchiveReason.Type)
		}
	}
}
