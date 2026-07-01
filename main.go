package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Model struct {
	ID string `json:"id"`
}

type ModelsResponse struct {
	Data []Model `json:"data"`
}

func main() {
	apiKey := os.Getenv("OPENROUTER_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "Error: OPENROUTER_API_KEY environment variable is required")
		os.Exit(1)
	}

	_, err := exec.LookPath("claude")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: 'claude' command not found in PATH")
		os.Exit(1)
	}

	models, err := fetchFreeModels()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching models: %v\n", err)
		os.Exit(1)
	}

	if len(models) == 0 {
		fmt.Fprintln(os.Stderr, "Error: No free models found")
		os.Exit(1)
	}

	printModels(models)

	selection := getUserSelection(len(models))
	selectedModel := models[selection-1].ID

	env := buildEnv(apiKey, selectedModel)

	cmd := exec.Command("claude")
	cmd.Env = env
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error launching claude: %v\n", err)
		os.Exit(1)
	}
}

func fetchFreeModels() ([]Model, error) {
	resp, err := http.Get("https://openrouter.ai/api/v1/models")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var result ModelsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	var freeModels []Model
	for _, m := range result.Data {
		if strings.HasSuffix(m.ID, ":free") {
			freeModels = append(freeModels, m)
		}
	}

	return freeModels, nil
}

func printModels(models []Model) {
	fmt.Println("Free models available:")
	for i, m := range models {
		fmt.Printf("%d. %s\n", i+1, m.ID)
	}
	fmt.Println()
}

func getUserSelection(max int) int {
	fmt.Printf("Select model (1-%d): ", max)
	var input string
	fmt.Scanln(&input)

	n, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil || n < 1 || n > max {
		fmt.Fprintf(os.Stderr, "Error: Invalid selection. Enter a number between 1 and %d\n", max)
		os.Exit(1)
	}
	return n
}

func buildEnv(apiKey, modelID string) []string {
	env := os.Environ()
	env = setEnv(env, "ANTHROPIC_BASE_URL", "https://openrouter.ai/api")
	env = setEnv(env, "ANTHROPIC_AUTH_TOKEN", apiKey)
	env = setEnv(env, "ANTHROPIC_API_KEY", "")
	env = setEnv(env, "ANTHROPIC_DEFAULT_OPUS_MODEL", modelID)
	env = setEnv(env, "ANTHROPIC_DEFAULT_SONNET_MODEL", modelID)
	env = setEnv(env, "ANTHROPIC_DEFAULT_HAIKU_MODEL", modelID)
	env = setEnv(env, "CLAUDE_CODE_SUBAGENT_MODEL", modelID)
	return env
}

func setEnv(env []string, key, value string) []string {
	kv := key + "=" + value
	for i, e := range env {
		if strings.HasPrefix(e, key+"=") {
			env[i] = kv
			return env
		}
	}
	return append(env, kv)
}