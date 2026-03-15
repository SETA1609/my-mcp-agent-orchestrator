package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os/exec"
	"testing"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
)

func TestHTTPIntegration(t *testing.T) {
	// Build the binary
	cmd := exec.Command("go", "build", "-o", "server", "./cmd/server")
	if err := cmd.Run(); err != nil {
		t.Fatalf("failed to build server: %v", err)
	}

	// For HTTP, we can use httptest.Server but since it's a binary, we'll start it on a port
	// For simplicity, assume it runs on localhost:8080 or use a random port

	// Start server in HTTP mode
	serverCmd := exec.Command("./server")
	serverCmd.Env = append(serverCmd.Env, "MCP_TRANSPORT=http", "HOST=localhost", "PORT=8080")

	if err := serverCmd.Start(); err != nil {
		t.Fatalf("failed to start server: %v", err)
	}

	// Wait for server to start
	time.Sleep(2 * time.Second)

	// Test /health endpoint
	resp, err := http.Get("http://localhost:8080/health")
	if err != nil {
		t.Fatalf("failed to get health: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
	resp.Body.Close()

	// For MCP over HTTP, we need to establish SSE connection and send messages
	// This is more complex. For now, test that the server starts and health works.

	// To test MCP, we need to:
	// 1. Establish SSE connection to /sse
	// 2. Send POST to /message with JSON-RPC

	// This requires more setup. For MVP, just test health.

	// Clean up
	if err := serverCmd.Process.Kill(); err != nil {
		t.Logf("failed to kill server: %v", err)
	}
	serverCmd.Wait()
}
