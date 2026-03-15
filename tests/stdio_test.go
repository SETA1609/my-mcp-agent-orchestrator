package tests

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
)

func TestStdioIntegration(t *testing.T) {
	// Build the binary
	cmd := exec.Command("go", "build", "-o", "server", "./cmd/server")
	if err := cmd.Run(); err != nil {
		t.Fatalf("failed to build server: %v", err)
	}

	// Start the server in stdio mode
	serverCmd := exec.Command("./server")
	serverCmd.Env = append(serverCmd.Env, "MCP_TRANSPORT=stdio")

	stdin, err := serverCmd.StdinPipe()
	if err != nil {
		t.Fatalf("failed to get stdin pipe: %v", err)
	}

	stdout, err := serverCmd.StdoutPipe()
	if err != nil {
		t.Fatalf("failed to get stdout pipe: %v", err)
	}

	if err := serverCmd.Start(); err != nil {
		t.Fatalf("failed to start server: %v", err)
	}

	reader := bufio.NewReader(stdout)

	// Read the initialize request or wait for server to be ready
	// MCP protocol starts with client sending initialize

	// For simplicity, send a tools/list request
	listRequest := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "tools/list",
		"params":  map[string]interface{}{},
	}

	sendRequest(t, stdin, listRequest)

	// Read response
	response := readResponse(t, reader)

	// Check if tools include echo
	if tools, ok := response["result"].(map[string]interface{})["tools"].([]interface{}); ok {
		foundEcho := false
		for _, tool := range tools {
			if toolMap, ok := tool.(map[string]interface{}); ok {
				if toolMap["name"] == "echo" {
					foundEcho = true
					break
				}
			}
		}
		if !foundEcho {
			t.Error("echo tool not found in tools/list response")
		}
	}

	// Send echo call
	echoRequest := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      2,
		"method":  "tools/call",
		"params": map[string]interface{}{
			"name": "echo",
			"arguments": map[string]interface{}{
				"input": "test message",
			},
		},
	}

	sendRequest(t, stdin, echoRequest)

	// Read response
	response = readResponse(t, reader)

	// Check result
	if result, ok := response["result"].(map[string]interface{}); ok {
		if content, ok := result["content"].([]interface{}); ok && len(content) > 0 {
			if textContent, ok := content[0].(map[string]interface{}); ok {
				if text, ok := textContent["text"].(string); ok {
					if text != "test message" {
						t.Errorf("expected 'test message', got %q", text)
					}
				}
			}
		}
	}

	// Clean up
	stdin.Close()
	if err := serverCmd.Wait(); err != nil {
		t.Logf("server exited with error: %v", err)
	}
}

func sendRequest(t *testing.T, writer io.Writer, req interface{}) {
	data, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("failed to marshal request: %v", err)
	}
	data = append(data, '\n')
	if _, err := writer.Write(data); err != nil {
		t.Fatalf("failed to write request: %v", err)
	}
}

func readResponse(t *testing.T, reader *bufio.Reader) map[string]interface{} {
	line, err := reader.ReadString('\n')
	if err != nil {
		t.Fatalf("failed to read response: %v", err)
	}
	line = strings.TrimSpace(line)
	var response map[string]interface{}
	if err := json.Unmarshal([]byte(line), &response); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}
	return response
}
