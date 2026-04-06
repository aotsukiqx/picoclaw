# pkg/providers

## OVERVIEW
30+ LLM provider implementations: OpenAI, Anthropic, Google Gemini, Azure, AWS Bedrock, Ollama, vLLM, and more.

## WHERE TO LOOK
| Task | Location | Notes |
|------|----------|-------|
| Provider factory | `factory.go` | Creates providers from config |
| Common interface | `types.go` | `Provider`, `Message`, `Delta` interfaces |
| HTTP base | `http_provider.go` | Base implementation for REST providers |
| Rate limiting | `ratelimiter.go` | Token bucket rate limiter |
| Error classification | `error_classifier.go` | Retry logic for provider errors |
| OpenAI compat | `openai_compat/` | OpenAI-compatible API providers |
| Anthropic | `anthropic/` | Anthropic Claude API |
| Azure | `azure/` | Azure OpenAI Service |
| AWS Bedrock | `bedrock/` | Bedrock with automatic region resolution |

## CONVENTIONS
- Provider `Name()` returns unique string (e.g., `"openai"`, `"anthropic"`)
- `Send(ctx, msgs) (Stream, error)` — core method signature
- `Stream` yields `Delta` structs with role, content, tool calls
- Error classification: `Retryable()`, `AuthError()`, `ContextExceeded()`

## ANTI-PATTERNS
- **DO NOT** hardcode API endpoints — use `api_base` from config
- **DO NOT** store API keys in provider code — use credential store
- **DO NOT** block on rate limiter — use non-blocking retry with backoff
