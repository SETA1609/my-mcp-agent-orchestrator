# PROJECT_TREE

> **Note:** Keep this file in sync with the actual directory layout. Run the `update-docs` workflow after any structural changes.

## Layout

<!-- Replace with your actual project structure. Example below: -->

```
<project-root>/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── resources/
│   │   ├── health.go
│   │   └── resources.go
│   ├── server/
│   │   └── server.go
│   ├── tools/
│   │   ├── echo.go
│   │   └── tools.go
│   └── transport/
│       ├── sse/
│       │   └── sse.go
│       ├── stdio/
│       │   └── stdio.go
│       └── transport.go
├── docs/
│   ├── agents/
│   │   ├── CLAUDE.md
│   │   ├── CODEX.md
│   │   ├── COPILOT.md
│   │   ├── GEMINI.md
│   │   └── OPENCODE.md
│   ├── INDEX.md
│   ├── PROJECT_TREE.md
│   ├── DEPENDENCIES.md
│   ├── CODE_STYLE.md
│   └── COMMIT_STYLE.md
├── .agents/
│   └── workflows/
│       ├── build.md
│       ├── commit.md
│       ├── review.md
│       └── update-docs.md
├── .env.example
├── .golangci.yml
├── .pre-commit-config.yaml
├── AGENTS.md
├── PLAN.md
├── go.mod
├── go.sum
└── README.md
```

## Key Directories

| Path                  | Purpose                              |
|-----------------------|--------------------------------------|
| `cmd/server/`         | Application entrypoint               |
| `internal/`           | Internal server, transport, tools    |
| `docs/`               | Project documentation                |
| `.agents/workflows/`  | Reusable agent workflow scripts      |
