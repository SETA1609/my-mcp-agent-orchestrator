// Package server constructs and configures the MCP server instance.
//
// # Registration pattern (confirmed against mcp-go v0.45)
//
// Capabilities attach to *server.MCPServer via:
//
//	s.AddTool(tool mcp.Tool, handler server.ToolHandlerFunc)
//	s.AddResource(resource mcp.Resource, handler server.ResourceHandlerFunc)
//
// Handler signatures:
//
//	ToolHandlerFunc     = func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
//	ResourceHandlerFunc = func(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error)
//
// Each capability module registers itself via a Register(*server.MCPServer) func.
// tools.RegisterAll / resources.RegisterAll call every module's Register in turn.
// To add a new capability:
//  1. Create internal/tools/<name>.go or internal/resources/<name>.go
//  2. Implement Register(s *server.MCPServer)
//  3. Add the call to tools.RegisterAll or resources.RegisterAll
//  4. No changes to this file are needed.
package server

import (
	"github.com/SETA1609/my-mcp-agent-orchestrator/internal/resources"
	"github.com/SETA1609/my-mcp-agent-orchestrator/internal/tools"
	mcpserver "github.com/mark3labs/mcp-go/server"
)

// New creates a configured MCPServer with all capabilities registered.
func New() *mcpserver.MCPServer {
	s := mcpserver.NewMCPServer("mcp-server", "0.1.0")
	tools.RegisterAll(s)
	resources.RegisterAll(s)
	return s
}
