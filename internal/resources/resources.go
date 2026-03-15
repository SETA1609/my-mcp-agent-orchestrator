// Package resources registers all MCP resources with the server.
//
// To add a new resource:
//  1. Create internal/resources/<name>.go
//  2. Implement Register(s *server.MCPServer) calling s.AddResource(definition, handler)
//  3. Add RegisterName(s) to RegisterAll below — no other files need to change.
package resources

import mcpserver "github.com/mark3labs/mcp-go/server"

// RegisterAll adds every resource in this package to s.
func RegisterAll(s *mcpserver.MCPServer) {
	RegisterHealth(s)
}
