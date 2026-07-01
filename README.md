# Claude Free Launcher

Claude Free is a launcher to easily use Claude Code with free models from OpenRouter.

<img width="847" height="529" alt="ezgif-55ba8d92fcef7eb2" src="https://github.com/user-attachments/assets/d3808bf8-7694-4aea-ada8-f3a6e49214ef" />

## Overview

This program acts as a bridge between Claude Code and OpenRouter's free models. It fetches available free models from OpenRouter, lets you select one, configures the necessary environment variables, and launches Claude Code with that model.

## Prerequisites

- Go 1.21+ (for building from source)
- `OPENROUTER_API_KEY` environment variable set (how to create the key [https://openrouter.ai/docs/api/reference/authentication](https://openrouter.ai/docs/api/reference/authentication))
- `claude` command available in PATH (Claude Code CLI)

## Installation

### Pre-built Binaries

Download the appropriate binary for your platform:
- Linux: `claude-free`
- Windows: `claude-free.exe`
- macOS (Intel): `claude-free`
- macOS (Apple Silicon): `claude-free`

Make it executable (Linux/macOS):
```bash
chmod +x claude-free
```

### Build from Source

```bash
git clone https://github.com/eutobias/claude-free
cd claude-free
go build -o claude-free main.go
```

## Configuration

Set your OpenRouter API key as an environment variable:

```bash
export OPENROUTER_API_KEY="your-api-key-here"
```

Add to your shell profile (`.bashrc`, `.zshrc`, etc.) for persistence.

## Usage

Run the launcher:

```bash
./claude-free
```

On Windows:
```cmd
claude-free.exe
```

The program will:
1. Verify `OPENROUTER_API_KEY` is set
2. Fetch free models from OpenRouter API
3. Display numbered list of models (those with `:free` suffix)
4. Prompt you to select a model by number
5. Set required environment variables
6. Launch Claude Code

## Environment Variables Set

The launcher configures these environment variables for the Claude Code process:

| Variable | Value |
|----------|-------|
| `ANTHROPIC_BASE_URL` | `https://openrouter.ai/api` |
| `ANTHROPIC_AUTH_TOKEN` | Value of `OPENROUTER_API_KEY` |
| `ANTHROPIC_API_KEY` | Empty string |
| `ANTHROPIC_DEFAULT_OPUS_MODEL` | Selected model ID |
| `ANTHROPIC_DEFAULT_SONNET_MODEL` | Selected model ID |
| `ANTHROPIC_DEFAULT_HAIKU_MODEL` | Selected model ID |
| `CLAUDE_CODE_SUBAGENT_MODEL` | Selected model ID |

## Error Handling

- Exits with code 1 if `OPENROUTER_API_KEY` not set
- Exits with code 1 if `claude` command not found in PATH
- Exits with code 1 if API request fails
- Exits with code 1 if no free models available
- Exits with code 1 if invalid selection entered

## How It Works

1. **API Call**: Fetches `https://openrouter.ai/api/v1/models`
2. **Filtering**: Selects models where `id` ends with `:free`
3. **Selection**: Interactive prompt for user choice
4. **Environment**: Builds modified environment inheriting parent vars plus overrides
5. **Exec**: Replaces process with `claude` command using new environment

## Credits
- To my friend [@MichelRibeiro](https://github.com/michelribeiro) who had this ideia
