# DEPENDENCIES

> **Note:** Keep this file current when adding, removing, or upgrading dependencies.

## Runtime Dependencies

| Package | Version | Purpose |
|---------|---------|---------|
| `github.com/mark3labs/mcp-go` | `v0.45.0` | MCP server SDK for stdio and HTTP/SSE transports |

## Development Dependencies

| Package | Version | Purpose |
|---------|---------|---------|
| `golangci-lint` | managed externally | Static analysis and formatting checks (`gofmt`, `govet`, `staticcheck`) |
| `pre-commit` | managed externally | Runs local commit hooks (`gofmt`, `golangci-lint`) |

## Upgrade Policy

- **Patch** versions: upgrade freely
- **Minor** versions: upgrade after reviewing the changelog
- **Major** versions: requires `[D]` (human) approval before upgrading
