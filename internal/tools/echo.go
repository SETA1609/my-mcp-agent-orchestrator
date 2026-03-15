// Echo tool — implement by [P|GF].
//
// Contract:
//   - RegisterEcho(s) calls s.AddTool(definition, handler)
//   - definition: mcp.NewTool("echo", mcp.WithDescription(...), mcp.WithString("input", ...))
//   - handler: func(ctx, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
//     extract req.Params.Arguments["input"], return mcp.NewToolResultText(input)
package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
)

// RegisterEcho adds the echo tool to s.
func RegisterEcho(s *mcpserver.MCPServer) {
	tool := mcp.NewTool("echo",
		mcp.WithDescription("Echoes the input string back"),
		mcp.WithString("input",
			mcp.Required(),
			mcp.Description("The string to echo"),
		),
	)
	s.AddTool(tool, echoHandler)
}

func echoHandler(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// [P|GF]: extract input and return it
	input, err := req.RequireString("input")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return mcp.NewToolResultText(input), nil
}
