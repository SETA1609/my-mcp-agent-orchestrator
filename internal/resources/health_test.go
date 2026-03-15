package resources

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
)

func TestHealthHandler(t *testing.T) {
	req := mcp.ReadResourceRequest{
		Params: struct {
			URI string `json:"uri"`
		}{
			URI: healthURI,
		},
	}

	contents, err := healthHandler(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(contents) != 1 {
		t.Fatalf("expected 1 content, got %d", len(contents))
	}

	textContent, ok := contents[0].(mcp.TextResourceContents)
	if !ok {
		t.Fatalf("expected TextResourceContents, got %T", contents[0])
	}

	if textContent.URI != healthURI {
		t.Errorf("expected URI %q, got %q", healthURI, textContent.URI)
	}

	if textContent.MIMEType != "application/json" {
		t.Errorf("expected MIMEType %q, got %q", "application/json", textContent.MIMEType)
	}

	// Parse the JSON
	var health map[string]string
	if err := json.Unmarshal([]byte(textContent.Text), &health); err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}

	expected := map[string]string{"status": "ok"}
	if health["status"] != expected["status"] {
		t.Errorf("expected status %q, got %q", expected["status"], health["status"])
	}
}
