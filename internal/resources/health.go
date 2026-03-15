// Health resource — implement by [P|GF].
//
// Contract:
//   - RegisterHealth(s) calls s.AddResource(definition, handler)
//   - definition: mcp.NewResource("mcp://health", "Health Status", ...)
//   - handler: func(ctx, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error)
//     return mcp.NewTextResourceContents(uri, mimeType, jsonPayload)
package resources

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
)

const healthURI = "mcp://health"

// RegisterHealth adds the health resource to s.
func RegisterHealth(s *mcpserver.MCPServer) {
	resource := mcp.NewResource(healthURI,
		"Health Status",
		mcp.WithResourceDescription("Server health status"),
		mcp.WithMIMEType("application/json"),
	)
	s.AddResource(resource, healthHandler)
}

func healthHandler(_ context.Context, _ mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	// [P|GF]: return JSON health payload
	return []mcp.ResourceContents{
		mcp.TextResourceContents{
			URI:      healthURI,
			MIMEType: "application/json",
			Text:     `{"status":"ok"}`,
		},
	}, nil
}
