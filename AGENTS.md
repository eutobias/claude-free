# AGENTS.md - Claude Free Launcher

## Project Overview

A Go CLI that launches Claude Code with free OpenRouter models. Fetches models from `https://openrouter.ai/api/v1/models`, filters for `:free` suffix, prompts user selection, sets required env vars, execs `claude`.

## Build

Output binary name: `claude-free` (no platform suffix)

```bash
# Linux
go build -o claude-free main.go

# Windows
GOOS=windows GOARCH=amd64 go build -o claude-free.exe main.go

# macOS (Intel)
GOOS=darwin GOARCH=amd64 go build -o claude-free main.go

# macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o claude-free main.go
```

## Run

```bash
export OPENROUTER_API_KEY="your-key"
./claude-free
```

## Required Environment

- `OPENROUTER_API_KEY` - Must be set before running

## Env Vars Set for Claude Code

| Variable | Value |
|----------|-------|
| ANTHROPIC_BASE_URL | https://openrouter.ai/api |
| ANTHROPIC_AUTH_TOKEN | $OPENROUTER_API_KEY |
| ANTHROPIC_API_KEY | (empty) |
| ANTHROPIC_DEFAULT_OPUS_MODEL | Selected model ID |
| ANTHROPIC_DEFAULT_SONNET_MODEL | Selected model ID |
| ANTHROPIC_DEFAULT_HAIKU_MODEL | Selected model ID |
| CLAUDE_CODE_SUBAGENT_MODEL | Selected model ID |

## Code Structure

Single file: `main.go`
- `fetchFreeModels()` - HTTP GET + JSON decode + filter `:free`
- `printModels()` - Numbered list
- `getUserSelection()` - Scan input, validate 1..N
- `buildEnv()` - Inherit os.Environ(), override 7 vars
- `main()` - Orchestration

## Testing

No tests currently. Manual test:
1. Set `OPENROUTER_API_KEY`
2. Run binary
3. Select model
4. Verify Claude Code launches

## Dependencies

Stdlib only: `net/http`, `encoding/json`, `os`, `os/exec`, `fmt`, `strings`, `strconv`