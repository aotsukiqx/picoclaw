# pkg/agent

## OVERVIEW
Core agent logic: context management, event bus, hooks, steering, subturn coordination, and the main agent loop.

## WHERE TO LOOK
| Task | File | Notes |
|------|------|-------|
| Agent loop | `loop.go` | Main agent execution loop |
| Context | `context.go`, `context_manager.go` | Request context lifecycle |
| EventBus | `eventbus.go` | Pub/sub for agent events |
| Hooks | `hooks.go`, `hook_process.go` | Pre/post processing hooks |
| Steering | `steering.go` | Inject messages mid-loop |
| SubTurn | `subturn.go` | Subagent coordination |
| Memory | `memory.go` | Agent memory interface |
| Model routing | `model_resolution.go` | Model selection logic |
| Thinking | `thinking.go` | Chain-of-thought handling |

## CONVENTIONS
- Accept interfaces, return structs (core design)
- EventBus for loose coupling between components
- Context carries request-scoped data through the loop
- Hooks: `PreTurn`, `PostTurn`, `PreTool`, `PostTool`, `PreSpawn`, `PostSpawn`

## ANTI-PATTERNS
- **DO NOT** mutate context after `Context.Prepare()` — treat as immutable
- **DO NOT** call `EventBus.Emit` synchronously in hot paths — use goroutines
- **DO NOT** block the agent loop — use async patterns for long operations
