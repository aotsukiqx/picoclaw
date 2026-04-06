# pkg/seahorse

## OVERVIEW
DAG-structured SQLite memory engine with FTS5 CJK support for agent short-term memory. Unique to PicoClaw.

## WHERE TO LOOK
| Task | File | Notes |
|------|------|-------|
| Engine core | `short_engine.go` | Main memory engine |
| Storage | `store.go` | SQLite persistence layer |
| Assembler | `short_assembler.go` | Assembles context from memory |
| Retrieval | `short_retrieval.go` | Similarity search with scoring |
| Compaction | `short_compaction.go` | Memory compression under pressure |
| Tool expansion | `tool_expand.go` | Expand tool call context |
| Schema | `schema.go` | SQLite schema definitions |

## CONVENTIONS
- DAG nodes: `MemoryNode{ID, Type, Content, Score, Created, Expires, Links}`
- FTS5 CJK tokenizer: `fts5 unicode61` for multilingual support
- Score range: 0.0-1.0 (relevance to current context)
- Compaction triggers when total size exceeds `MaxMemoryBudget`

## ANTI-PATTERNS
- **DO NOT** call `Store.Save()` on every write — batch writes
- **DO NOT** query FTS5 without score filtering — always use `score > threshold`
- **DO NOT** store raw tool output in memory — summarize first
