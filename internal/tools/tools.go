// Package tools registers all MCP tools with the server.
//
// To add a new tool:
//  1. Create internal/tools/<name>.go
//  2. Implement Register(s *server.MCPServer) calling s.AddTool(definition, handler)
//  3. Add RegisterName(s) to RegisterAll below — no other files need to change.
package tools

import mcpserver "github.com/mark3labs/mcp-go/server"

// RegisterAll adds every tool in this package to s.
func RegisterAll(s *mcpserver.MCPServer) {
	RegisterEcho(s)
}
