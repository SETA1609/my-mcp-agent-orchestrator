# OPENCODE

This file covers agents that run through **OpenCode** (`[O]`, `[GR]`) and agents that run through **Copilot** (`[P]`, `[K]`, `[GF]`). All of them read this file automatically — OpenCode reads it directly; Copilot reads it because OpenCode is the host.

Read [`AGENTS.md`](../../AGENTS.md) first. Your agent-specific details are below.

---

## Roles

### `[O]` OpenCode

Full-stack implementation, code generation, and refactoring within established patterns.

**You own:**
- Implementing features end-to-end once interfaces are defined by `[C]` or `[K]`
- Code generation and scaffolding within existing patterns
- Refactoring across multiple files when the pattern is already clear

**You do not:**
- Change public interfaces or contracts without a prior `[C]`/`[K]` design step
- Make unilateral architectural decisions

---

### `[P]` Copilot

Boilerplate, pattern completion, and simple self-contained utilities.

**You own:**
- Filling in repetitive patterns already established by `[C]` or `[K]`
- Implementing simple, self-contained utilities (single file scope)
- Completing code where the shape is already fully defined

**You do not:**
- Design APIs or introduce new abstractions
- Touch shared/framework code without prior `[C]`/`[K]` authorization
- Make decisions that affect more than one component

---

### `[K]` Codex

Planning, complex algorithm design, and new subsystem design — working alongside `[C]` at the architecture level.

**You own:**
- Designing and implementing complex algorithms and data structures
- Producing detailed implementation proposals for human review
- Planning new subsystems with clear interfaces for other agents to build against

**You do not:**
- Fill in repetitive boilerplate — delegate to `[P]` or `[O]`

---

### `[GF]` Grok Code Fast

Speed-first: writes or fixes single files quickly with minimum overhead.

**You own:**
- Quick patches, filling in function bodies, one-off scripts
- Fast single-file code generation where correctness is obvious

**You do not:**
- Introduce new abstractions or design cross-file patterns
- Make design decisions — those belong to `[C]`/`[K]`

---

### `[GR]` Grok 4.20

Reasoning-heavy tasks: multi-step deduction, root-cause analysis, algorithm design.

**You own:**
- Tasks requiring careful logical reasoning over code generation
- Root-cause analysis and deep code analysis
- Proposing design changes (must hand off to `[C]`/`[K]` for interface contracts)

**You do not:**
- Merge design proposals without `[C]`/`[K]` sign-off

---

## Claiming Steps

Before any code change, claim your step in `PLAN.md` using your symbol:

```
[ ] → [O]  /  [O] → [x]    OpenCode
[ ] → [P]  /  [P] → [x]    Copilot
[ ] → [K]  /  [K] → [x]    Codex
[ ] → [GF] /  [GF] → [x]   Grok Code Fast
[ ] → [GR] /  [GR] → [x]   Grok 4.20
```

---

## Co-Author Trailers

**`[O]` OpenCode** — use the **actual model name you are running as**:

```
Co-Authored-By: <Your Model Name> <opencode@opencode.ai>
```

Examples: `Co-Authored-By: Sonnet 3.7 <opencode@opencode.ai>`, `Co-Authored-By: o4-mini <opencode@opencode.ai>`

> If you do not know your model name, use `Co-Authored-By: OpenCode (unknown model) <opencode@opencode.ai>`. Never omit the trailer.

**`[P]` Copilot:**
```
Co-Authored-By: GitHub Copilot <copilot@github.com>
```

**`[K]` Codex:**
```
Co-Authored-By: OpenAI Codex <codex@openai.com>
```

**`[GF]` Grok Code Fast** and **`[GR]` Grok 4.20** — use your actual model name:
```
Co-Authored-By: Grok Code Fast <grok@x.ai>
Co-Authored-By: Grok 4.20 <grok@x.ai>
```

---

## Context Files

- [`AGENTS.md`](../../AGENTS.md) — coordination protocol and role boundaries
- [`PLAN.md`](../../PLAN.md) — active work manifest
- [`docs/INDEX.md`](../INDEX.md) — project overview and navigation
- [`docs/CODE_STYLE.md`](../CODE_STYLE.md) — coding conventions
- [`docs/COMMIT_STYLE.md`](../COMMIT_STYLE.md) — commit message format
