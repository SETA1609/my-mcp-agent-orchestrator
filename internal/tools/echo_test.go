package tools

import (
	"context"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
)

func TestEchoHandler(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		hasError bool
	}{
		{
			name:     "echo simple string",
			input:    "hello",
			expected: "hello",
			hasError: false,
		},
		{
			name:     "echo empty string",
			input:    "",
			expected: "",
			hasError: false,
		},
		{
			name:     "echo with spaces",
			input:    "hello world",
			expected: "hello world",
			hasError: false,
		},
		{
			name:     "missing input",
			input:    "", // not used
			expected: "",
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := map[string]any{
				"input": tt.input,
			}
			if tt.name == "missing input" {
				args = map[string]any{}
			}

			req := mcp.CallToolRequest{
				Params: struct {
					Name      string         `json:"name"`
					Arguments map[string]any `json:"arguments"`
				}{
					Name:      "echo",
					Arguments: args,
				},
			}

			result, err := echoHandler(context.Background(), req)

			if tt.hasError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if len(result.Content) != 1 {
				t.Errorf("expected 1 content item, got %d", len(result.Content))
				return
			}

			textContent, ok := result.Content[0].(mcp.TextContent)
			if !ok {
				t.Errorf("expected TextContent, got %T", result.Content[0])
				return
			}

			if textContent.Text != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, textContent.Text)
			}
		})
	}
}
