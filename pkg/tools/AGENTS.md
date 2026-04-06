# pkg/tools

## OVERVIEW
Tool implementations for PicoClaw: I2C, SPI, exec, cron, filesystem, shell, web, spawn, skills, MCP, and message handling.

## WHERE TO LOOK
| Task | File | Notes |
|------|------|-------|
| Tool registry | `registry.go` | Tool registration and discovery |
| Cron/scheduling | `cron.go` | Cron job management |
| Shell/exec | `shell.go`, `shell_process_*.go` | Command execution |
| Spawn/async | `spawn.go` | Subagent spawning |
| Skills | `skills_install.go`, `skills_search.go` | Skill management |
| MCP | `mcp_tool.go` | MCP protocol tool |
| Filesystem | `filesystem.go` | File operations |
| Web search | `search_tool.go` | Web search integration |
| Session | `session.go` | Session state management |

## CONVENTIONS
- Platform-specific code: `*_unix.go`, `*_windows.go` build tags
- Tool implementations return `ToolResult` (see `result.go`)
- Error handling via `result.Error(err)` helper
- Async operations use goroutines with context cancellation

## ANTI-PATTERNS
- **DO NOT** use blocking I/O in tool callbacks — use goroutines
- **DO NOT** store sensitive data in tool results — filter before returning
